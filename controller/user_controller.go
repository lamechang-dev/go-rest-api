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

type UserController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &UserController{uu}
}

func (uc *UserController) SignUp(c echo.Context) error {
	return nil
}

func (uc *UserController) Login(c echo.Context) error {
	return nil
}

func (uc *UserController) Logout(c echo.Context) error {
	return nil
}