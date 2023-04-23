package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed: ", err)
	}
}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed: ", err)
	}

	h := Holdings{
		Amount:        1,
		Purchased:     time.Now(),
		PurchasePrice: 19,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed: ", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id sent back: ", err)
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed: ", err)
	}

	h := Holdings{
		Amount:        1,
		Purchased:     time.Now(),
		PurchasePrice: 19,
	}
	if _, err = testRepo.InsertHolding(h); err != nil {
		t.Error("insert failed: ", err)
	}

	result, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("insert failed: ", err)
	}

	if len(result) != 1 {
		t.Error("invalid number of results: got ", len(result))
	}
}

func TestSQLiteRepository(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed: ", err)
	}
}
