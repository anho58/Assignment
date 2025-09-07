package http

import (
	"net/http"
	"strings"

	application "github.com/anho58/Assignment/part1_Golang/api_gateway/internal/application"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandler struct {
	service application.BookServiceInterface
}

func NewBookHandler(service application.BookServiceInterface) *BookHandler {
	return &BookHandler{service: service}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, role, err := ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// store user info in context
		c.Set("userID", userID)
		c.Set("role", role)
		c.Next()
	}
}

func AuthorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (h *BookHandler) RegisterRoutes(rg *gin.RouterGroup) {
	books := rg.Group("/books")
	books.Use(AuthMiddleware())
	{
		books.GET("", h.List)
		books.POST("", h.Add)
		books.GET(":id", h.Get)
		books.DELETE(":id", AuthorMiddleware(), h.Delete)
	}
}

// GetAll godoc
// @Summary List all books
// @Tags books
// @Produce json
// @Security BearerAuth
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

// Add godoc
// @Summary add a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body application.CreateBookDTO true "Book info"
// @Security BearerAuth
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

// Get godoc
// @Summary Get a book by ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Security BearerAuth
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
// @Security BearerAuth
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
