package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koga456/sample-api/model/entity"
)

type CocktailRepository interface {
	GetCocktails() (cocktails []entity.CocktailEntity, err error)
}

type cocktailRepository struct{}

func NewCocktailRepository() CocktailRepository {
	return &cocktailRepository{}
}

func (tr *cocktailRepository) GetCocktails() (
	cocktails []entity.CocktailEntity, err error) {
	cocktails = []entity.CocktailEntity{}
	rows, err := Db.
		Query("SELECT id, name, description FROM cocktail ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		cocktail := entity.CocktailEntity{}
		err = rows.Scan(&cocktail.ID, &cocktail.Name, &cocktail.Description)
		if err != nil {
			log.Print(err)
			return
		}
		cocktails = append(cocktails, cocktail)
	}

	return
}
