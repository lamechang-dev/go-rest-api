package controller

import (
	"go-rest-api/usecase"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context)	error
	Login(c echo.Context)	error
	Logout(c echo.Context)	error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	return nil
}

func (uc *userController) Login(c echo.Context) error {
	return nil
}

func (uc *userController) Logout(c echo.Context) error {
	return nil
}