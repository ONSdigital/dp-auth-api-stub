package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/ONSdigital/dp-auth-api-stub/handler"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/dp-auth-api-stub/config"
	"github.com/gorilla/mux"
	"github.com/ONSdigital/go-ns/server"
)

func main() {
	log.HumanReadable = true
	log.Namespace = "dp-auth-api-stub"

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	errorChan := make(chan error, 1)

	cfg, err := config.Get()
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}

	log.Info("config on startup", log.Data{"config": cfg})
	apiStub, err := handler.NewAPIStub()

	if err != nil {
		log.ErrorC("failed to create apiStub, exiting", err, nil)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Path("/identity").Methods(http.MethodGet).HandlerFunc(apiStub.Identify)
	router.Path("/service").Methods(http.MethodPost).HandlerFunc(apiStub.CreateServiceAccount)

	httpServer := server.New(cfg.BindAddr, router)

	go func() {
		log.Info("starting http server", nil)
		if err := httpServer.ListenAndServe(); err != nil {
			errorChan <- err
		}
	}()

	// wait until fatal event
	select {
	case err := <-errorChan:
		log.ErrorC("application error encountered, commencing graceful shutdown", err, nil)
	case sig := <-signals:
		log.Info("received os signal, commencing graceful shutting down", log.Data{"signal": sig.String()})
	}

	// shutdown
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	if err := httpServer.Shutdown(ctx); err != nil {
		log.ErrorC("graceful shutdown failed", err, nil)
	} else {
		log.Info("graceful shutdown completed without error", nil)
	}
}
