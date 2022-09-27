package controller

import "net/http"

type Router interface {
	//r? w? interfaceとtypeとstruct
	HandleCocktailsRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc CocktailController
}

func NewRouter(tc CocktailController) Router {
	return &router{tc}
}

func (ro *router) HandleCocktailsRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetCocktails(w, r)
	default:
		w.WriteHeader(405)
	}
}
