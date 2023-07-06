package cars

import (
	"github.com/gofiber/fiber/v2"
)

// GetCars func gets all exists cars.
// @Description Get all exists cars.
// @Summary get all exists cars
// @Tags Cars
// @Accept json
// @Produce json
// @Success 200 {array} Car
// @Router /cars [get]
func (h handler) GetCars(c *fiber.Ctx) error {
	var items []Car

	if result := h.DB.Find(&items); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&items)
}
