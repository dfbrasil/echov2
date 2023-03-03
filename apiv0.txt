package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct{}

type BooksParams struct{
	Offset int `query:"offset"`
	Limit int `query:"limit"`
}

type BookIDParams struct{
	ID int `param:"id"`
}

type PostBook struct{
	Title string `json:"title"`
}

var (
	books = []string{"Livro 1", "Livro 2", "Livro 3"}
)

func (a *API) getBooks(c echo.Context)  error {
	
	params := &BooksParams{}

	err := c.Bind(params)
	if err != nil{ //se tem erro devolve o status de bad request
		return c.JSON(http.StatusBadRequest, "Parâmetro de query inválidos") //c.jason devolve um status e um erro
	}

	if params.Offset > len(books) || params.Offset < 0 {
		return c.JSON(http.StatusBadRequest, "Parâmetro de query inválidos")
	}

	if params.Limit < 0 || params.Limit > len(books) {
		return c.JSON(http.StatusBadRequest, "Parâmetro de query inválidos")
	}

	var from, to int

	if params.Offset > 0 {
		from = params.Offset
	}

	if params.Limit > 0{
		to = params.Limit
	} else {
		to = len(books)
	}

	return c.JSON(http.StatusOK, books[from:to])
}

func (a *API) getBook(c echo.Context) error{
	
	params := &BookIDParams{}
	
	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Parâmentros inválidos")
	}

	index := params.ID - 1

	if index < 0 || index > len(books)-1 {
		return c.JSON(http.StatusBadRequest, "Invalid Parametros")
	}

	return c.JSON(http.StatusOK, books[index])
}

func (a *API) postBook(c echo.Context) error{
	book := &PostBook{}

	err := c.Bind(book)
	if err != nil{
		return c.JSON(http.StatusBadRequest, "Invalid parametro")
	}

	books = append(books, book.Title)
	return c.NoContent(http.StatusCreated)
}