package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/models"
	"gorm.io/gorm"
)

// CreateConditions seeds db with Condition-entries
func CreateConditions(db *gorm.DB, conditions []models.Condition) error {

	for _, condition := range conditions {
		if err := db.Create(condition).Error; err != nil {
			return err
		}
	}
	return nil
}
