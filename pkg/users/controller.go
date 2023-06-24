package users

import (
	"github.com/edwinpaye/gots/pkg/midleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/users")
	routes.Post("/", midleware.JWTProtected(), h.AddUser)
	routes.Get("/", midleware.JWTProtected(), h.GetUsers)
	routes.Get("/:id", midleware.JWTProtected(), h.GetUser)
	// routes.Put("/:id", midleware.JWTProtected(), h.UpdateUser)
	routes.Delete("/:id", midleware.JWTProtected(), h.DeleteUser)
}
