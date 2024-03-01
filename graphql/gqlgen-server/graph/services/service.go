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
	IssueService
}

type services struct {
	*userService
	*repoService
	*projectService
	*issueService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:    &userService{exec: exec},
		repoService:    &repoService{exec: exec},
		projectService: &projectService{exec: exec},
		issueService:   &issueService{exec: exec},
	}
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
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

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
}
