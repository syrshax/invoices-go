package main

import (
	"github.com/syrshax/invoice-go-v2/handlers"
	"github.com/syrshax/invoice-go-v2/server"
	"log"
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
	app.server.AddHandler("/jobs/", handlers.JobStatus)
	app.server.AddHandler("/download/", handlers.Download)
	app.server.Run()
}
