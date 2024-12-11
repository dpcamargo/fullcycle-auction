package user_controller

import (
	"context"
	"net/http"

	"github.com/dpcamargo/fullcycle-auction/configuration/rest_err"
	"github.com/dpcamargo/fullcycle-auction/internal/usecase/user_usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userController struct {
	userUseCase user_usecase.UserUseCase
}

func NewUserController(userUsecase user_usecase.UserUseCase) *userController {
	return &userController{
		userUseCase: userUsecase,
	}
}

func (u *userController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "invalid UUID value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUseCase.FindUserByID(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}
