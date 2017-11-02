package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/allisonverdam/best-credit-card/app"
	"github.com/allisonverdam/best-credit-card/controllers"
	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/errors"
	"github.com/allisonverdam/best-credit-card/services"
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
	address := map[bool]string{true: os.Getenv("PORT"), false: strconv.Itoa(app.Config.ServerPort)}[os.Getenv("PORT") != ""]
	logger.Infof("Server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(fmt.Sprintf(":%s", address), nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/handlers
		return c.Write("pong!")
	})

	/**
	* @apiDefine ContentTypeJson
	* @apiHeaderExample {json} Headers:
	*     {
	*		"Content-Type": "application/json"
	*     }
	**/
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

	/**
	* @apiDefine AuthRequired
	* @apiHeaderExample {json} Headers:
	*     {
	*       "Authorization": "Bearer {{TOKEN}}",
	*		"Content-Type": "application/json"
	*     }
	**/
	//Fazendo as requisições desse group passarem pelo middleware de auth
	rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
		TokenHandler: services.JWTHandler,
	}))

	//Instanciando DAOs
	cardDAO := daos.NewCardDAO()
	userDAO := daos.NewPersonDAO()
	walletDAO := daos.NewWalletDAO()

	//Fazendo o load dos resources
	controllers.ServeCardResource(rg, services.NewCardService(cardDAO))
	controllers.ServePersonResource(rg, services.NewPersonService(userDAO))
	controllers.ServeWalletResource(rg, services.NewWalletService(walletDAO))

	//Ignorar o middleware de auth
	controllers.ServeAuthResource(router.Group("/v1"), services.NewAuthService(userDAO))

	return router
}
