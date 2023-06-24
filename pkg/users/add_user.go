package users

import (
	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AddUserRequestBody struct {
	Name     string `json:"title"`
	Lastname string `json:"author"`
	Age      string `json:"description"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateBook func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body AddUserRequestBody true "User Data"
// @Success 201 {object} models.User
// @Security ApiKeyAuth
// @Router /users [post]
func (h handler) AddUser(c *fiber.Ctx) error {
	body := AddUserRequestBody{}

	// parse body, attach to AddUserRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var entity models.User

	entity.Name = body.Name
	entity.Lastname = body.Lastname
	entity.Age = body.Age
	entity.Email = body.Email
	// entity.Password = generatePassword(body.Password)
	err2 := generatePassword2(&entity, body.Password)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}

	// insert new db entry
	if result := h.DB.Create(&entity); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&entity)
}

// func generatePassword(raw string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "error generating password", err
// 	}
// 	return string(hash), nil
// }

func generatePassword2(u *models.User, raw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
