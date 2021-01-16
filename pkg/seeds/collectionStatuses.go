package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/models"
	"gorm.io/gorm"
)

// CreateCollectionStatuses seeds db with CollectionStatus-entries
func CreateCollectionStatuses(db *gorm.DB, statuses []models.CollectionStatus) error {

	for _, status := range statuses {
		if err := db.Create(status).Error; err != nil {
			return err
		}
	}
	return nil
}
