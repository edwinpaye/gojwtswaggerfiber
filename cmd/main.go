package main

import (
	"log"

	// swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/edwinpaye/gots/docs"
	"github.com/edwinpaye/gots/pkg/books"
	"github.com/edwinpaye/gots/pkg/common/config"
	"github.com/edwinpaye/gots/pkg/common/db"
	"github.com/edwinpaye/gots/pkg/jwt"
	"github.com/edwinpaye/gots/pkg/midleware"
	"github.com/edwinpaye/gots/pkg/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title API documentation
// @version 1.0.0
// @host localhost:3000
// @BasePath /
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email test@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	midleware.FiberMiddleware(app)
	// app.Group("/swagger")
	// app.Get("*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.HandlerDefault)
	// app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL:         "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:3080/swagger/oauth2-redirect.html",
	// }))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	books.RegisterRoutes(app, db)
	users.RegisterRoutes(app, db)
	jwt.TokenRoutes(app)

	app.Listen(c.Port)

}
