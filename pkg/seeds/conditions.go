package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/models"
	"gorm.io/gorm"
)

// CreateCondition seeds db with Condition-entry
func CreateCondition(db *gorm.DB, names []string) error {

	for _, name := range names {
		if err := db.Create(&models.Condition{Name: name}).Error; err != nil {
			return err
		}
	}
	return nil
}
