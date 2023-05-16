package api

import (
	"example/model"
	"example/model/response"
	"example/services/user_service/repository"
	"log"
	"net/http"

	"example/controllers/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase *user.UserUsecase
}

func NewUserHandler() UserHandlerInterface {
	userRepo := repository.NewUserRepository()
	log.Println("User Repo is initialized")
	return &UserHandler{
		userUsecase: user.NewUserUseCase(userRepo),
	}
}

type UserHandlerInterface interface {
	GetUser(context *gin.Context)
}

// Healthcheck API Handler
func (hc UserHandler) GetUser(context *gin.Context) {
	user, err := hc.userUsecase.GetUser(context, model.UserInput{
		UserRole: "role",
	})
	if err != nil {
		log.Print("Error : %v", err)
	}
	context.JSON(http.StatusOK,
		response.NewHTTPResponse(200, "OK", user))

}
