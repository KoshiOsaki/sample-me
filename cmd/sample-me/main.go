package main

import (
	"fmt"
	"net/http"

	"github.com/koga456/sample-api/controller"
	"github.com/koga456/sample-api/model/repository"
)

var tr = repository.NewCocktailRepository()
var tc = controller.NewCocktailController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/cocktails/", ro.HandleCocktailsRequest)
	fmt.Println("Server is running at localhost:8080")
	server.ListenAndServe()
}
