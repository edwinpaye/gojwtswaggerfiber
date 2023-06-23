package main

import (
	"log"

	// swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/swagger"
	// _ "github.com/arsmn/fiber-swagger/v2/example/docs"
	_ "github.com/edwinpaye/gots/docs"
	"github.com/edwinpaye/gots/pkg/books"
	"github.com/edwinpaye/gots/pkg/common/config"
	"github.com/edwinpaye/gots/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	// app.Group("/swagger")
	// app.Get("*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)

}
