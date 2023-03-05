package api

import (

	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) { //registra as rotas da API
	//(a *API é um receiver)
	//RegisterRoutes recebe um router de ECHO

	e.Use(requestIDHandler)

	//Criação de grupo de rotas públicas e privadas
	public := e.Group("")
	protected := e.Group("", authMidlleware)

	//Rotas
	public.GET("/books", a.getBooks)
	public.GET("/books/:id", a.getBook)
	protected.POST("/books", a.postBook)
	public.DELETE("/books/:id", a.deleteBook)
	public.PUT("/books/:id", a.updateBook)
	public.GET("/:name", a.Parametros)
}
