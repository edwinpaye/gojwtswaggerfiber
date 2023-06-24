package users

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// GetBook func gets user by given ID or 404 error.
// @Description Get user by given ID.
// @Summary get user by given ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (h handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var body models.User

	if result := h.DB.First(&body, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&body)
}
