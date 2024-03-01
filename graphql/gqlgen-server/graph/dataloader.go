package graph

import (
	"context"
	"errors"
	"graphql-github-sample/graph/model"
	"graphql-github-sample/graph/services"

	"github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(srv services.Services) *Loaders {
	userBatcher := &userBatcher{
		Srv: srv,
	}

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](userBatcher.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (u *userBatcher) BatchGetUsers(
	ctx context.Context, IDs []string,
) []*dataloader.Result[*model.User] {
	results := make([]*dataloader.Result[*model.User], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}

	indexs := make(map[string]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	users, err := u.Srv.ListUsersByID(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for _, user := range users {
		var rsl *dataloader.Result[*model.User]
		if err != nil {
			rsl = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexs[user.ID]] = rsl
	}
	return results
}
