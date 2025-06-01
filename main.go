package main

import (
	"log"

	"github.com/syrshax/invoice-go-v2/handlers"
	"github.com/syrshax/invoice-go-v2/server"
)

type Application struct {
	server *server.Server
}

func main() {
	log.Printf("Welcome to CLI debug to invoice-go-v2! \n")
	app := Application{
		server: server.NewServer("8000"),
	}

	app.server.AddHandler("/", handlers.Home)
	app.server.AddHandler("/upload", handlers.Upload)
	app.server.Run()
}
