package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/allisonverdam/go-api-mcc/app"
	"github.com/allisonverdam/go-api-mcc/controllers"
	"github.com/allisonverdam/go-api-mcc/daos"
	"github.com/allisonverdam/go-api-mcc/errors"
	"github.com/allisonverdam/go-api-mcc/services"
	dbx "github.com/go-ozzo/ozzo-dbx"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	_ "github.com/lib/pq"
)

func main() {
	// Carrega as configurações
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Configuração inválida: %s", err))
	}

	// Carrega as mensagens de erro
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// Cria o logger
	logger := logrus.New()

	// Conecta ao banco de dados
	db, err := dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("Server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("/v1")

	rg.Post("/auth", controllers.Auth(app.Config.JWTSigningKey))
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		SigningMethod: app.Config.JWTSigningMethod,
		TokenHandler:  controllers.JWTHandler,
	}))

	// Fazendo o load dos controllers
	cardDAO := daos.NewCardDAO()
	uerDAO := daos.NewPersonDAO()
	controllers.ServeCardResource(rg, services.NewCardService(cardDAO))
	controllers.ServePersonResource(router.Group("/v1"), services.NewPersonService(uerDAO))

	return router
}