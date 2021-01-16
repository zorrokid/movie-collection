package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/models"
	"gorm.io/gorm"
)

// CreateCaseTypes seeds db with CaseType-entries
func CreateCaseTypes(db *gorm.DB, caseTypes []models.CaseType) error {

	for _, caseType := range caseTypes {
		if err := db.Create(caseType).Error; err != nil {
			return err
		}
	}
	return nil
}
