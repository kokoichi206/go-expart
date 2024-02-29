package services

import (
	"context"
	"graphql-github-sample/graph/db"
	"graphql-github-sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repoService struct {
	exec boil.ContextExecutor
}

func (r *repoService) GetRepoByID(ctx context.Context, id string) (*model.Repository, error) {
	repo, err := db.FindRepository(ctx, r.exec, id,
		db.RepositoryColumns.ID, db.RepositoryColumns.Name, db.RepositoryColumns.Owner, db.RepositoryColumns.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &model.Repository{
		ID:   repo.ID,
		Name: repo.Name,
		Owner: &model.User{
			ID: repo.Owner,
		},
		CreatedAt: repo.CreatedAt,
	}, nil
}

func (r *repoService) GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)

	if err != nil {
		return nil, err
	}

	return &model.Repository{
		ID:   repo.ID,
		Name: repo.Name,
		Owner: &model.User{
			ID: repo.Owner,
		},
		CreatedAt: repo.CreatedAt,
	}, nil
}
