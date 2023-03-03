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

// type PostBook struct{
// 	Title string `json:"title"`
// }

var (
	books = map[int]*book{}
	seq = 1
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

	return c.JSON(http.StatusOK, books)
}

func (a *API) getBook(c echo.Context) error {

	params := &BookIDParams{}
	
	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Parâmentros inválidos")
	}

	index := params.ID - 1

	if index < 0 || index > len(books)-1 {
		return c.JSON(http.StatusBadRequest, "Parâmentros inválidos")
	}
	
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func (a *API) postBook(c echo.Context) error {
	
	b := &book{
		ID: seq,
	}

	if err := c.Bind(b); err != nil{
		return c.JSON(http.StatusBadRequest, "Parâmetro Inválido")
	}

	books[b.ID] = b
	seq++
	return c.NoContent(http.StatusCreated)
}

func (a *API) deleteBook(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	delete(books,id)
	return c.NoContent(http.StatusNoContent)
}

func (a *API) updateBook(c echo.Context) error {
	b := new(book)
	if err := c.Bind(b); err != nil{
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Title = b.Title
	return c.JSON(http.StatusOK, books[id])
}