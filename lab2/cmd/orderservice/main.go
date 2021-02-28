package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	orderservice "lab2/pkg/orderservice"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("./log/my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	killSignalChan := getKillSignalChan()
	srv := startServer(":8000")
	waitForKillSignal(killSignalChan)
	log.Fatal(srv.Shutdown(context.Background()))
}

func getKillSignalChan() chan os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return signalChan
}

func waitForKillSignal(signalChan <-chan os.Signal) {
	killSignal := <-signalChan

	switch killSignal {
	case os.Interrupt:
	case syscall.SIGINT:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}

func startServer(serverURL string) *http.Server {
	log.WithFields(log.Fields{"url": serverURL}).Info("starting the server")
	router := orderservice.Router()

	srv := &http.Server{Addr: serverURL, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}
