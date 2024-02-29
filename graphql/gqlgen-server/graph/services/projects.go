package services

import (
	"context"
	"fmt"
	"graphql-github-sample/graph/db"
	"graphql-github-sample/graph/model"
	"net/url"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type projectService struct {
	exec boil.ContextExecutor
}

func (p *projectService) GetProjectByID(ctx context.Context, id string) (*model.ProjectV2, error) {
	project, err := db.FindProject(ctx, p.exec, id,
		db.ProjectColumns.ID,
		db.ProjectColumns.Title,
		db.ProjectColumns.Number,
		db.ProjectColumns.URL,
		db.ProjectColumns.Owner,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find project: %w", err)
	}

	url, err := url.Parse(project.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse project URL: %w", err)
	}

	return &model.ProjectV2{
		ID:     project.ID,
		Title:  project.Title,
		Number: int(project.Number),
		URL: model.MyURL{
			URL: *url,
		},
		Owner: &model.User{
			ID: project.Owner,
		},
	}, nil
}

func (p *projectService) GetProjectByOwnerAndNumber(ctx context.Context, ownerID string, number int) (*model.ProjectV2, error) {
	project, err := db.Projects(
		qm.Select(
			db.ProjectColumns.ID,
			db.ProjectColumns.Title,
			db.ProjectColumns.Number,
			db.ProjectColumns.URL,
			db.ProjectColumns.Owner,
		),
		db.ProjectWhere.Owner.EQ(ownerID),
		db.ProjectWhere.Number.EQ(int64(number)),
	).One(ctx, p.exec)

	if err != nil {
		return nil, fmt.Errorf("failed to find project: %w", err)
	}

	url, err := url.Parse(project.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse project URL: %w", err)
	}

	return &model.ProjectV2{
		ID:     project.ID,
		Title:  project.Title,
		Number: int(project.Number),
		URL: model.MyURL{
			URL: *url,
		},
		Owner: &model.User{
			ID: project.Owner,
		},
	}, nil
}
