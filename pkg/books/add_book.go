package books

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Books
// @Accept json
// @Produce json
// Param title body string true "Title"
// Param author body string true "Author"
// @Param book body AddBookRequestBody true "Book Data"
// Param book_attrs body models.BookAttrs true "Book attributes"
// @Success 201 {object} models.Book
// @Security JWT
// @Router /books [post]
func (h handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestBody{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	// insert new db entry
	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&book)
}
