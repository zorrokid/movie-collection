package main

import (
	"log"

	"github.com/zorrokid/movieCollection/pkg/database"
	"github.com/zorrokid/movieCollection/pkg/seeds"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("failed connecting")
	}

	database.MigrateDB(db)

	for _, seed := range seeds.All() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
