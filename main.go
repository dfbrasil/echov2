package main

import (
	"echov2/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleIndex(c echo.Context) error{ //as funções em Echo devem retornar um error
	return c.JSON(http.StatusOK, map[string]string{"message":"hello!"})//echo possui uma função no context que se chama JSON, que recebe um status e um dado(interface)
}

func main(){
	e := echo.New() // um objeto de ECHO que está retornando um ponteiro de ECHO

	e.GET("/", handleIndex) //a este objeto Echo está mapeada uma rota

	a := &api.API{}
	a.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))//se ocorre algum error na inicialização será mostrado no terminal
}
