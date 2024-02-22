package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"soap-server/config"
	"soap-server/internal/services"
	soaphandlers "soap-server/pkg/transport/handlers"
	"time"
)

func main() {
	configs, err := config.GetConfigs(".env")
	if err != nil {
		log.Fatalln("error from config:", err)
	}

	srvs := services.InitServices(configs.ApiKey)

	srv := &http.Server{
		Addr:         configs.ServerAddress + configs.Port,
		Handler:      soaphandlers.InitSoapHandler(srvs),
		ReadTimeout:  10 * time.Millisecond,
		WriteTimeout: 10 * time.Millisecond,
	}

	go func() {
		log.Println("Starting server at ", configs.ServerAddress+configs.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = srv.Shutdown(ctx); err != nil {
		log.Println("Cannot shutdown gracefully:", err)
	}
	log.Println("Shutting down gracefully")
}
