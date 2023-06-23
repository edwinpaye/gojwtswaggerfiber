package jwt

import (
	"github.com/edwinpaye/gots/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// func for describe group of public routes.
func TokenRoutes(a *fiber.App) {

	route := a.Group("/token")

	route.Get("/new", GetNewAccessToken) // create a new access tokens
}

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
