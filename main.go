package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	//_ "github.com/your/repo/Docs"
	RoutesPing "github.com/your/repo/Controller/Ping"
	Produto "github.com/your/repo/Controller/Produtos"
)

func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "github.com/your/repo/swagger/swagger.json",
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
	app.Get("/Ping", RoutesPing.Ping)
	app.Get("/produto/", Produto.Produto)
	app.Get("/produto/:id", Produto.ProdutoById)
	app.Listen(":3000")
}
