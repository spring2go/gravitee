package cmd

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/spring2go/gravitee/log"
	"github.com/spring2go/gravitee/services"
	"github.com/urfave/negroni"
	graceful "gopkg.in/tylerb/graceful.v1"
)

// RunServer runs the app
func RunServer(configFile string) error {
	cfg, db, err := initConfigDB(configFile)
	if err != nil {
		return err
	}
	defer db.Close()

	// start the services
	if err := services.Init(cfg, db); err != nil {
		return err
	}
	defer services.Close()

	// Start a classic negroni app
	app := negroni.New()
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(negroni.NewStatic(http.Dir("public")))

	// Create a router instance
	router := mux.NewRouter()

	// Add routes
	services.HealthService.RegisterRoutes(router, "/v1")
	services.UserService.RegisterRoutes(router, "/v1/user")
	services.OauthService.RegisterRoutes(router, "/v1/oauth")
	services.WebService.RegisterRoutes(router, "/web")

	// Set the router
	app.UseHandler(router)

	log.INFO.Println("Starting gravitee server on port ", cfg.ServerPort)
	// Run the server on $ServerPort, gracefully stop on SIGTERM signal
	graceful.Run(":"+strconv.Itoa(cfg.ServerPort), 5*time.Second, app)

	return nil
}
