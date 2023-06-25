package books

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// DeleteBook func for deletes book by given ID.
// @Description Delete book by given ID.
// @Summary delete book by given ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 {} status "no content"
// @Security JWT
// @Router /books/{id} [delete]
func (h handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete book from db
	h.DB.Delete(&book)

	return c.SendStatus(fiber.StatusNoContent)
}
