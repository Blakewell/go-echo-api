package handler

import (
	"github.com/labstack/echo/v4"
	"go-echo-api/package/entity"
	"go-echo-api/package/user"
	"net/http"
)

type UserService struct {
	service user.Service
}

func MakeUserHandlers(e *echo.Echo, service *user.Service) UserService {

	var userService = UserService{
		service: *service,
	}

	e.GET("/users", userService.findAllUsers)
	e.POST("/users", userService.addUser)

	return userService
}

func (s *UserService) findAllUsers(context echo.Context) error {
	all, err := s.service.FindAll()
	if err == nil {
		return context.JSON(http.StatusOK, all)
	}
	return err
}

func (s *UserService) addUser(context echo.Context) error {

	u := new(entity.User)
	if err := context.Bind(u); err != nil {
		return err
	}

	uID, _ := s.service.Store(u)

	return context.JSON(http.StatusCreated, uID)

}
