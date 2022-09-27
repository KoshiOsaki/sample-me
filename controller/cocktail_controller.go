package controller

import (
	"encoding/json"
	"net/http"

	"github.com/koga456/sample-api/controller/dto"
	"github.com/koga456/sample-api/model/repository"
)

type CocktailController interface {
	GetCocktails(w http.ResponseWriter, r *http.Request)
}

type cocktailController struct {
	tr repository.CocktailRepository
}

//不明
func NewCocktailController(tr repository.CocktailRepository) CocktailController {
	return &cocktailController{tr}
}

func (tc *cocktailController) GetCocktails(w http.ResponseWriter, r *http.Request) {
	cocktails, err := tc.tr.GetCocktails()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var cocktailResponses []dto.CocktailResponse
	for _, v := range cocktails {
		//IDは？
		cocktailResponses = append(cocktailResponses, dto.CocktailResponse{Name: v.Name, Description: v.Description})
	}

	var cocktailsResponse dto.CocktailsResponse
	cocktailsResponse.Cocktails = cocktailResponses

	output, _ := json.MarshalIndent(cocktailsResponse.Cocktails, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}
