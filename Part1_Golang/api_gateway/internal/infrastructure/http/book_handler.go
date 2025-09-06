package http

import (
	"net/http"

	application "github.com/anho58/Assignment/Part1_Golang/api_gateway/internal/application"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandler struct {
	service application.BookServiceInterface
}

func NewBookHandler(service application.BookServiceInterface) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) RegisterRoutes(rg *gin.RouterGroup) {
	books := rg.Group("/books")
	{
		books.GET("", h.List)
		books.POST("", h.Add)
		books.GET(":id", h.Get)
		books.DELETE(":id", h.Delete)
	}
}

// GetAll godoc
// @Summary List all books
// @Tags books
// @Produce json
// @Router /books [get]
func (h *BookHandler) List(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := h.service.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Create godoc
// @Summary Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body application.CreateBookDTO true "Book info"
// @Router /books [post]
func (h *BookHandler) Add(c *gin.Context) {
	var req application.CreateBookDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Add(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetByID godoc
// @Summary Get a book by ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Router /books/{id} [get]
func (h *BookHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete a book by ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Router /books/{id} [delete]
func (h *BookHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
