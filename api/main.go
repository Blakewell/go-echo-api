package main

import (
	"github.com/labstack/echo/v4"
	"go-echo-api/api/handler"
	"go-echo-api/package/entity"
	"go-echo-api/package/user"
)

func main() {

	e := echo.New()

	repo := user.NewInMemRepository()
	service := user.NewService(repo)

	u1 := &entity.User{Name: "Blake"}
	u2 := &entity.User{Name: "Kacee"}

	_, _ = service.Store(u1)
	_, _ = service.Store(u2)

	handler.MakeUserHandlers(e, service)

	e.Logger.Fatal(e.Start(":1323"))

}
