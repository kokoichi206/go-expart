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

type issueService struct {
	exec boil.ContextExecutor
}

func (i *issueService) GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error) {
	issue, err := db.Issues(
		qm.Select(
			db.IssueColumns.ID,
			db.IssueColumns.URL,
			db.IssueColumns.Title,
			db.IssueColumns.Closed,
			db.IssueColumns.Number,
			db.IssueColumns.Author,
			db.IssueColumns.Repository,
		),
		db.IssueWhere.Repository.EQ(repoID),
		db.IssueWhere.Number.EQ(int64(number)),
	).One(ctx, i.exec)

	if err != nil {
		return nil, fmt.Errorf("failed to find issue: %w", err)
	}

	url, err := url.Parse(issue.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse issue URL: %w", err)
	}

	return &model.Issue{
		ID: issue.ID,
		URL: model.MyURL{
			URL: *url,
		},
		Title:  issue.Title,
		Closed: (issue.Closed == 1),
		Number: int(issue.Number),
		Author: &model.User{
			ID: issue.Author,
		},
		Repository: &model.Repository{
			ID: issue.Repository,
		},
	}, nil
}
