package users

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// GetUsers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Security ApiKeyAuth
// @Router /users [get]
func (h handler) GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}
