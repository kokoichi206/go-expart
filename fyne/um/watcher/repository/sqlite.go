package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	// real means float64
	query := `
		CREATE TABLE IF NOT EXISTS holdings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			amount REAL NOT NULL,
			purchased INTEGER NOT NULL,
			purchase_price INTEGER NOT NULL
		);
	`

	_, err := r.Conn.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}

func (r *SQLiteRepository) InsertHolding(h Holdings) (*Holdings, error) {
	query := `
		INSERT INTO holdings (amount, purchased, purchase_price)
		VALUES (?, ?, ?)
	`

	// stmt, err := r.Conn.Prepare(query)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to prepare a statement: %w", err)
	// }
	// defer stmt.Close()

	// res, err := stmt.Exec(h.Amount, h.Purchased, h.PurchasePrice)

	resp, err := r.Conn.Exec(query, h.Amount, h.Purchased.Unix(), h.PurchasePrice)
	if err != nil {
		return nil, fmt.Errorf("failed to execute a statement: %w", err)
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last inserted id: %w", err)
	}

	h.ID = id

	return &h, nil
}

func (r *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := `
		SELECT id, amount, purchased, purchase_price
		FROM holdings
		ORDER BY purchased
	`

	rows, err := r.Conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute a query: %w", err)
	}
	defer rows.Close()

	holdings := []Holdings{}

	for rows.Next() {
		h := Holdings{}
		var unixTime int64

		err := rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan a row: %w", err)
		}

		h.Purchased = time.Unix(unixTime, 0)

		holdings = append(holdings, h)
	}

	return holdings, nil
}

func (r *SQLiteRepository) GetHoldingByID(id int64) (*Holdings, error) {
	row := r.Conn.QueryRow(`
		SELECT id, amount, purchased, purchase_price
		FROM holdings WHERE id = ?;
		`)

	var h Holdings
	var unixTime int64
	if err := row.Scan(
		&h.ID,
		&h.Amount,
		&unixTime,
		&h.PurchasePrice,
	); err != nil {
		return nil, fmt.Errorf("failed to scan a row: %w", err)
	}

	h.Purchased = time.Unix(unixTime, 0)

	return &h, nil
}

func (r *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	if id == 0 {
		return fmt.Errorf("id is invalid")
	}

	stmt := `
		UPDATE holdings
		SET amount = ?, purchased = ?, purchase_price = ?
		WHERE id = ?;
	`

	resp, err := r.Conn.Exec(stmt, updated.Amount, updated.Purchased.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return fmt.Errorf("failed to execute a statement: %w", err)
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

func (r *SQLiteRepository) DeleteHolding(id int64) error {
	resp, err := r.Conn.Exec("DELETE FROM holdings WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to execute a statement: %w", err)
	}

	ra, err := resp.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if ra == 0 {
		return errDeleteFailed
	}

	return nil
}
