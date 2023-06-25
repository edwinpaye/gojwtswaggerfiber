package jwt

import (
	"encoding/json"
	"log"

	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/edwinpaye/gots/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type SignInDTO struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=3, max=100"`
}

type Email struct {
	email string
}

// @Description SignIn a user.
// @Summary signIn a user
// @Tags Token
// @Accept json
// @Produce json
// @Param user body SignInDTO true "Credentials Data"
// @Success 200 {object} object
// @Router /token [post]
func (h handler) SignIn(c *fiber.Ctx) error {
	body := SignInDTO{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var entity models.User
	// email := &Email{email: body.Email}
	var user = models.User{Email: body.Email}
	res2B, _ := json.Marshal(body)
	log.Default().Println(string(res2B))

	// var password = body.Password
	if result := h.DB.First(&entity, user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if err2 := body.ValidatePassword(entity.Password); err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// return c.Status(fiber.StatusOK).JSON(&entity)
	// Generate a new Access token.
	token, err3 := utils.GenerateNewAccessToken()
	if err3 != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err3.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
		"user":         entity,
	})
}

func (u *SignInDTO) ValidatePassword(p string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(u.Password))
	// err := bcrypt.CompareHashAndPassword([]byte(p), []byte(u.Password))
	// if err != nil {
	// 	return
	// }
	// return nil
}
