package database

import (
	"fmt"

	"github.com/zorrokid/movieCollection/pkg/models"
)

// GetDbConstants reads collects constant kind of values from database to a DbConstants object
func GetDbConstants() *DbConstants {

	constants := &DbConstants{}

	db, err := InitDB()
	if err != nil {
		panic("failed connecting")
	}

	condition := models.Condition{}

	db.First(&condition)
	fmt.Printf("Got %s", condition.Name)

	var conditions []models.Condition

	constants.ConditionIds = make(map[string]uint)
	db.Find(&conditions)
	for _, c := range conditions {
		fmt.Printf("Got condition %s with id %d\n", c.Name, c.ID)
		constants.ConditionIds[c.Name] = c.ID
	}

	var statuses []models.CollectionStatus

	constants.StatusIds = make(map[string]uint)

	db.Find(&statuses)
	for _, s := range statuses {
		fmt.Printf("Got status %s with id %d\n", s.Name, s.ID)
		constants.StatusIds[s.Name] = s.ID
	}
	return constants
}
