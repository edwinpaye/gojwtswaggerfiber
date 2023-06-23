package books

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// UpdateBook func for updates book by given ID.
// @Description Update book.
// @Summary update book
// @Tags Books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body UpdateBookRequestBody true "Book Data"
// Param title body string true "Title"
// Param author body string true "Author"
// Param book_status body integer true "Book status"
// Param book_attrs body models.BookAttrs true "Book attributes"
// @Success 200 {object} models.Book
// @Security ApiKeyAuth
// @Router /books/{id} [put]
func (h handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateBookRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	// save book
	h.DB.Save(&book)

	return c.Status(fiber.StatusOK).JSON(&book)
}
