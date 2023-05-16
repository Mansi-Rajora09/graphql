package repository

import (
	"context"
	"example/model"
)

type User interface {
	FetchAll(ctx context.Context, search string) ([]model.UserInfo, error)
}
