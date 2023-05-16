package user

import (
	"context"
	"errors"
	"example/model"
	"example/model/response"
	"example/services/user_service/repository"
	"net/http"
	"strings"

	"github.com/vektah/gqlparser/gqlerror"
	"go.uber.org/zap"
)

type UserRunner interface {
	CreateUser(ctx context.Context, input model.UserInput) (*model.User, error)
	GetUser(ctx context.Context, request model.UserInput) (interface{}, error)
}
type UserUsecase struct {
	repo repository.User
}

func NewUserUseCase(repo repository.User) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (oc *UserUsecase) GetUser(ctx context.Context, request model.UserInput) (interface{}, error) {
	if len(strings.TrimSpace(request.UserRole)) == 0 {
		return response.NewHTTPResponse(http.StatusBadRequest, "invalid params", nil), errors.New("invalid params")
	}
	// Call the FetchAll method to get all the users
	response, err := oc.repo.FetchAll(ctx, request.UserRole)
	zap.S().Infow("Get User Request",
		zap.Any("request", request),
		zap.Any("response", response),
		zap.Error(err),
	)
	if err != nil {
		return nil, gqlerror.Errorf("Error while parsing response ")
	}
	return response, nil
}
