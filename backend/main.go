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
	bukuRepo := repository.NewBukuRepository(db)

	mainAPI := api.NewAPI(*usersRepo, *bukuRepo)
	mainAPI.Start()
}
