package main

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-32/backend/api"
	"github.com/rg-km/final-project-engineering-32/backend/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "backend/db/perpustakaan.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
	productsRepo := repository.NewProductRepository(db)
	cartItemRepo := repository.NewCartItemRepository(db)
	salesRepo := repository.NewSalesRepository(db)
	transactionRepo := repository.NewTransactionRepository(db, *productsRepo, *cartItemRepo)

	mainAPI := api.NewAPI(*usersRepo, *productsRepo, *cartItemRepo, transactionRepo, *salesRepo)
	mainAPI.Start()
}
