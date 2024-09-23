package delivery

import (
	"net/http"
	"project-golang-crud/domains"
	"strconv"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookUsecase domains.BookUsecase
}

func NewBookHandler(e *echo.Echo, bookUsecase domains.BookUsecase) {
	handler := &BookHandler{BookUsecase: bookUsecase}

	e.POST("/books", handler.Create)
	e.GET("/books", handler.GetAll)
	e.GET("/books/:id", handler.GetByID)
	e.PUT("/books/:id", handler.Update)
	e.DELETE("/books/:id", handler.Delete)
}
// Create godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept json
// @Produce json
// @Param book body domains.Book true "Book payload"
// @Success 201 {object} domains.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *BookHandler) Create(c echo.Context) error {
var book domains.Book
if err := c.Bind(&book); err != nil {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
}

if err := h.BookUsecase.Create(&book); err != nil {
	return c.JSON(http.StatusInternalServerError, err)
}

return c.JSON(http.StatusCreated, book)
}

// GetAll godoc
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags books
// @Produce json
// @Success 200 {array} domains.Book
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (h *BookHandler) GetAll(c echo.Context) error {
books, err := h.BookUsecase.GetAll()
if err != nil {
	return c.JSON(http.StatusInternalServerError, err)
}

return c.JSON(http.StatusOK, books)
}

// GetByID godoc
// @Summary Get book by ID
// @Description Get a single book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} domains.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [get]
func (h *BookHandler) GetByID(c echo.Context) error {
idStr := c.Param("id")
id, err := strconv.Atoi(idStr)
if err != nil {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
}

book, err := h.BookUsecase.GetByID(uint(id))
if err != nil {
	return c.JSON(http.StatusInternalServerError, err)
}

if book == nil {
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Book not found"})
}

return c.JSON(http.StatusOK, book)
}

// Update godoc
// @Summary Update a book
// @Description Update a book's information by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body domains.Book true "Book payload"
// @Success 200 {object} domains.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [put]
func (h *BookHandler) Update(c echo.Context) error {
idStr := c.Param("id")
id, err := strconv.Atoi(idStr)
if err != nil {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
}

var book domains.Book
if err := c.Bind(&book); err != nil {
	return c.JSON(http.StatusBadRequest, err)
}

book.ID = uint(id)

if err := h.BookUsecase.Update(&book); err != nil {
	return c.JSON(http.StatusInternalServerError, err)
}

return c.JSON(http.StatusOK, book)
}

// Delete godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [delete]
func (h *BookHandler) Delete(c echo.Context) error {
idStr := c.Param("id")
id, err := strconv.Atoi(idStr)
if err != nil {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
}

if err := h.BookUsecase.Delete(uint(id)); err != nil {
	return c.JSON(http.StatusInternalServerError, err)
}

return c.NoContent(http.StatusNoContent)
}
