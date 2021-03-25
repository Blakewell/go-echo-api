package main

import (
	"github.com/labstack/echo/v4"
	"go-echo-api/package/entity"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		u := entity.User{Name: "Blake"}
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
