package handler

import (
	"github.com/labstack/echo/v4"
	"go-echo-api/package/entity"
	"go-echo-api/package/user"
	"net/http"
)

func MakeUserHandlers(e *echo.Echo, service *user.Service) {

	e.GET("/users", func(context echo.Context) error {
		all, err := service.FindAll()
		if err == nil {
			return context.JSON(http.StatusOK, all)
		}
		return err
	})

	e.POST("/users", func(context echo.Context) error {

		u := new(entity.User)
		if err := context.Bind(u); err != nil {
			return err
		}

		uID, _ := service.Store(u)

		return context.JSON(http.StatusCreated, uID)

	})
}
