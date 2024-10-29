package handlers

import (
	"github.com/aryanicosa/golang-fullstack-lms/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandlersInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type UserHandlersImpl struct {
	userService service.UserServiceInterface
}

func NewUserHandlers(userService service.UserServiceInterface) UserHandlersInterface {
	return &UserHandlersImpl{
		userService: userService,
	}
}

func (h *UserHandlersImpl) Login(ctx *gin.Context) {

}

func (h *UserHandlersImpl) Register(ctx *gin.Context) {

}