package dto

type CocktailResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CocktailRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CocktailsResponse struct {
	Cocktails []CocktailResponse `json:"cocktails"`
}
