package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/models"
	"gorm.io/gorm"
)

// CreateCase seeds db with Case-entry
func CreateCase(db *gorm.DB, names []string) error {

	for _, name := range names {
		if err := db.Create(&models.Case{Name: name}).Error; err != nil {
			return err
		}
	}
	return nil
}
