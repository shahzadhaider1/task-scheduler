package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-openapi/loads"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	runtime "github.com/shahzadhaider1/task-scheduler"
	"github.com/shahzadhaider1/task-scheduler/config"
	"github.com/shahzadhaider1/task-scheduler/gen/restapi"
	"github.com/shahzadhaider1/task-scheduler/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}
	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	api := handlers.NewHandler(ctx, rt, swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	defer server.Shutdown()

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	done := make(chan bool)

	go gracefulShutdown(ctx, server, rt, done)

	if err := server.Serve(); err != nil {
		panic(err)
	}

	<-done
	log().Info("Server stopped gracefully")
}

func log() *logger.Entry {
	level, err := logger.ParseLevel(viper.GetString(config.LogLevel))
	if err != nil {
		logger.SetLevel(logger.DebugLevel)
	}
	logger.SetLevel(level)

	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "main",
	})
}

func gracefulShutdown(ctx context.Context, server *restapi.Server, rt *runtime.Runtime, done chan<- bool) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTSTP, syscall.SIGTERM)

	<-quit

	log().Info("server is shutting down...")

	if err := server.Shutdown(); err != nil {
		logger.Warnf("could not gracefully shutdown the server: %+v", err)
	}

	log().Info("Closing db connections")

	if err := rt.Service().Close(ctx); err != nil {
		logger.Warnf("could not gracefully shutdown the mongo client: %+v", err)
	}
	close(done)
}
