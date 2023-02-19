package main

import (
	"context"
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/products", productHandler)

	server := http.Server{
		Addr:         ":3030",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("server up and running... on http://localhost:3000")
		error := server.ListenAndServe()
		if error != nil {
			log.Fatal(error)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signal := <-signalChannel
	log.Println("Recieved terminate command, Graceful Shutdown.", signal)

	textContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(textContext)
}
