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
	// routes.Post("/", h.AddBook)
	routes.Get("/", midleware.JWTProtected(), h.GetUsers)
	// routes.Get("/:id", h.GetBook)
	// routes.Put("/:id", h.UpdateBook)
	// routes.Delete("/:id", h.DeleteBook)
}
