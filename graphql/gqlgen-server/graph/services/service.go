package services

import (
	"context"
	"graphql-github-sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	UserService
	RepoService
	ProjectService
}

type services struct {
	*userService
	*repoService
	*projectService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:    &userService{exec: exec},
		repoService:    &repoService{exec: exec},
		projectService: &projectService{exec: exec},
	}
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepoService interface {
	GetRepoByID(ctx context.Context, id string) (*model.Repository, error)
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type ProjectService interface {
	GetProjectByID(ctx context.Context, id string) (*model.ProjectV2, error)
	GetProjectByOwnerAndNumber(ctx context.Context, ownerID string, number int) (*model.ProjectV2, error)
}
