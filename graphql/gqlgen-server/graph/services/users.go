package services

import (
	"context"
	"fmt"
	"graphql-github-sample/graph/db"
	"graphql-github-sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func (u *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := db.FindUser(ctx, u.exec, id,
		db.UserTableColumns.ID, db.UserTableColumns.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec) // limit 1

	if err != nil {
		return nil, fmt.Errorf("failed to get user by name: %w", err)
	}

	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
