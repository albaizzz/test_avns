package actions

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"

	"test_avns/apitest/infrastructures/adapter"

	"test_avns/apitest/infrastructures"
	"test_avns/apitest/routes"
	config "github.com/spf13/viper"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

// RunServer serves the service
func RunServer(c *cli.Context) {
	// Set application configuration
	if err := infrastructures.SetConfiguration(c.String("config")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set log configuration
	infrastructures.SetLog()

	// Set router
	route := new(routes.Route)
	route.Database = routes.InitSQL()
	route.MySQLConnector = adapter.NewMySQLConnector(routes.SetupMySQL())
	route.Redis = routes.InitRedis()
	routeHandler := route.Init()
	n := negroni.New()
	recovery := negroni.NewRecovery()
	// If debug is false, do not show any panic stack or message related to error.
	if config.GetBool("app.debug") == false {
		recovery.PrintStack = false
		recovery.Formatter = &routes.StaticTextPanicFormatter{}
	}
	n.Use(recovery)
	n.Use(infrastructures.NewHttpLogger())
	n.UseHandler(routeHandler)

	s := &http.Server{
		Addr:         ":" + config.GetString("app.port"),
		Handler:      n,
		ReadTimeout:  time.Duration(config.GetInt("app.read_timeout")) * time.Second,
		WriteTimeout: time.Duration(config.GetInt("app.write_timeout")) * time.Second,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		logrus.Info("Shutting down server...")
		if err := s.Shutdown(context.Background()); err != nil {
			logrus.Errorf("http server got %#v", err)
		}
	}()

	// Listen and serve
	logrus.WithFields(map[string]interface{}{
		"component": "http",
		"port":      config.GetInt("app.port"),
	}).Info("starting test_avns/apitest server")

	err := s.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.Errorf("http server got %#v", err)
	}

	logrus.Info("Server gracefully stopped")

}
