package main

import (
	"fmt"
	"log"
	"net/http"

	"inMemoryDb/handler"
	"inMemoryDb/repository"
	"inMemoryDb/service"
)

func main() {
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	hndlr := handler.NewHandler(svc)

	http.HandleFunc("/create", hndlr.Create)
	http.HandleFunc("/get", hndlr.GetAll)

	port := ":8080"
	fmt.Println("Server http://localhost" + port + " running...")
	log.Fatal(http.ListenAndServe(port, nil))
}
