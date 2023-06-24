package users

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// DeleteUser func for deletes user by given ID.
// @Description Delete user by given ID.
// @Summary delete user by given ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {} status "no content"
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func (h handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var body models.User

	if result := h.DB.First(&body, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&body)

	return c.SendStatus(fiber.StatusNoContent)
}
