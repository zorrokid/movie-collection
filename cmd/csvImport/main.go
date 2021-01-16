package main

import (
	"fmt"

	"github.com/zorrokid/movieCollection/pkg/database"
	"github.com/zorrokid/movieCollection/pkg/models"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("failed connecting")
	}

	condition := models.Condition{}

	db.First(&condition)
	fmt.Printf("Got %s", condition.Name)

	var conditions []models.Condition

	db.Find(&conditions)
	for _, c := range conditions {
		fmt.Printf("Got %s", c.Name)
	}
}
