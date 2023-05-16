package repository

import (
	"context"
	"example/model"
)

type user struct {
}

func NewUserRepository() *user {
	return &user{}
}
func (p *user) FetchAll(ctx context.Context, search string) ([]model.UserInfo, error) {
	// Here we have hard Coded the value but we will query the data from the database instead
	result := []model.UserInfo{
		{
			ID:   1,
			Name: "Mansi",
		},
	}

	return result, nil
}
