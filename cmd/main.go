package main

import (
	"html/template"
	"log"

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
// @securityDefinitions.apikey JWT
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
	// app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		DefaultModelExpandDepth: 0,
		Filter:                  swagger.FilterConfig{Enabled: true},
		PersistAuthorization:    true,
		DocExpansion:            "none",
		CustomStyle: template.CSS(`
			h1,h2,h3,h4,h5,a,pre { color: #8c8cfa !important; text-color:white !important; }
			body { background-color: #000000; text-color: white; }
			div, span, { color: white !important; text-color: white !important; }
			.opblock, .models { background-color : gray !important; }
			.opblock-summary-method { border-color: black !important; }
		`),
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	books.RegisterRoutes(app, db)
	users.RegisterRoutes(app, db)
	jwt.TokenRoutes(app, db)

	app.Listen(c.Port)

}
