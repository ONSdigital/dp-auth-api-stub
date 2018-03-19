package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
	"github.com/ONSdigital/dp-auth-api-stub/handler"
)

func main() {
	log.HumanReadable = true
	log.Namespace = "dp-auth-api-stub"

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	errorChan := make(chan error, 1)

	apiStub, err := handler.NewAPIStub()
	if err != nil {
		log.ErrorC("failed to create apiStub, exiting", err, nil)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Path("/identity").Methods(http.MethodGet).HandlerFunc(apiStub.Identify)
	router.Path("/service").Methods(http.MethodPost).HandlerFunc(apiStub.CreateServiceAccount)

	httpServer := server.New(":8082", router)

	go func() {
		log.Info("starting http server", nil)
		if err := httpServer.ListenAndServe(); err != nil {
			errorChan <- err
		}
	}()

	shutdown := func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.ErrorC("graceful shutdown failed", err, nil)
		} else {
			log.Info("graceful shutdown completed without error", nil)
		}
	}

	select {
	case err := <-errorChan:
		log.ErrorC("application error encountered, commencing graceful shutdown", err, nil)
		shutdown()
	case sig := <-signals:
		log.Info("received os signal, commencing graceful shutting down", log.Data{"signal": sig.String()})
		shutdown()
	}

}
