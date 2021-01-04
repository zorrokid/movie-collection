package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/models"
	"gorm.io/gorm"
)

// CreateStatus seeds db with Case-entry
func CreateStatus(db *gorm.DB, names []string) error {

	for _, name := range names {
		if err := db.Create(&models.Status{Name: name}).Error; err != nil {
			return err
		}
	}
	return nil
}
