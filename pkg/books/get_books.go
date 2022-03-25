package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-api-medium/pkg/common/models"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.Status(http.StatusNotFound)

		return
	}

	c.JSON(http.StatusOK, &books)
}
