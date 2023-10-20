package app

import (
	"log"
	"net/http"

	"backend_task/internal/server"
	"backend_task/internal/service"
	"backend_task/internal/transport/rest"
)

func Run() {
	// TODO Init Router&Server
	router := http.NewServeMux()
	server := new(server.Server)

	// TODO Init Services
	clientService := service.NewServiceClient()
	parserService := service.NewServiceParser()

	// TODO Init Hadlers
	clientHandler := rest.NewHandlerClient(clientService)
	parserHandler := rest.NewHandlerParser(parserService)

	// TODO Call Handlers
	clientHandler.Init(router)
	parserHandler.Init(router)

	// TODO Server Run
	if err := server.Run("8080", router); err != nil {
		log.Fatalln("error with starting server")
		return
	}
}
