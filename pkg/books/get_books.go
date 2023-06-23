package books

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func (h handler) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}
