package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zorrokid/movieCollection/cmd/csvImport/loader"
	"github.com/zorrokid/movieCollection/pkg/database"
	"github.com/zorrokid/movieCollection/pkg/models"
)

func main() {

	if len(os.Args) == 1 {
		panic("No input filename argument")
	}

	filename := os.Args[1]

	err := loader.LoadCollectionItems(filename)

	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.InitDB()
	if err != nil {
		panic("failed connecting")
	}

	condition := models.Condition{}

	db.First(&condition)
	fmt.Printf("Got %s", condition.Name)

	var conditions []models.Condition

	conditionIds := make(map[string]uint)
	db.Find(&conditions)
	for _, c := range conditions {
		fmt.Printf("Got condition %s with id %d\n", c.Name, c.ID)
		conditionIds[c.Name] = c.ID
	}

	var statuses []models.CollectionStatus

	statusIds := make(map[string]uint)

	db.Find(&statuses)
	for _, s := range statuses {
		fmt.Printf("Got status %s with id %d\n", s.Name, s.ID)
		statusIds[s.Name] = s.ID
	}
}
