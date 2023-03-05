package api

import (

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type API struct{}

type (
	book struct {
		ID   int    `json:"id"`
		Title string `json:"title"`
	}
)

type BooksParams struct{
	Offset int `query:"offset"`
	Limit int `query:"limit"`
}

type BookIDParams struct{
	ID int `param:"id"`
}

type HelloWorld struct {
    Message string `json:"message"`
}

var (
	books = map[int]*book{}
	seq = 1
)

func (a *API) getBooks(c echo.Context)  error {
	
	params := &BooksParams{}

	err := c.Bind(params)
	if err != nil{ //se tem erro devolve o status de bad request
		return c.JSON(http.StatusBadRequest, "invalid query parameters") //c.jason devolve um status e um erro
	}

	if params.Offset > len(books) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, "invalid query parameters")
	}

	if params.Limit < 0 || params.Limit > len(books) {
		return c.JSON(http.StatusBadRequest, "invalid query parameters")
	}

	return c.JSON(http.StatusOK, books)
}

func (a *API) getBook(c echo.Context) error {

	params := &BookIDParams{}
	
	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid parameters")
	}

	index := params.ID - 1

	if index < 0 || index > len(books)-1 {
		return c.JSON(http.StatusBadRequest, "invalid parameters")
	}
	
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not convert parameter 'id' to integer"})
	}
	return c.JSON(http.StatusOK, books[id])
}

func (a *API) postBook(c echo.Context) error {
	
	b := &book{
		ID: seq,
	}

	if err := c.Bind(b); err != nil{
		return c.JSON(http.StatusBadRequest, "invalid parameter")
	}

	books[b.ID] = b
	seq++
	return c.NoContent(http.StatusCreated)
}

func (a *API) deleteBook(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid parameter")
	}

	delete(books,id)
	return c.NoContent(http.StatusNoContent)
}

func (a *API) updateBook(c echo.Context) error {

	b := new(book)

	if err := c.Bind(b); err != nil{
		return c.JSON(http.StatusBadRequest, "invalid parameter")
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not convert parameter 'id' to integer"})
	}

	books[id].Title = b.Title
	return c.JSON(http.StatusOK, books[id])
}

func (a *API) Parametros(c echo.Context) error {
    params, err := c.Param("name"), echo.ErrBadRequest

	if err != nil{
		return c.JSON(http.StatusOK, HelloWorld{
			Message: "Olá Mundo, meu nome é " + params, 
		})
	}
	
    return err
}
