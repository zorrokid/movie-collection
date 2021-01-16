package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/models"
	"github.com/zorrokid/movieCollection/pkg/seed"
	"gorm.io/gorm"
)

// All returns instances of seeds
func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateConditions",
			Run: func(db *gorm.DB) error {
				return CreateConditions(db, []models.Condition{
					models.NewCondition(1, "Bad - Damaged"),
					models.NewCondition(2, "Poor - Slightly Damaged"),
					models.NewCondition(3, "Fair"),
					models.NewCondition(4, "Good"),
					models.NewCondition(5, "Excellent"),
					models.NewCondition(6, "Still Wrapped"),
				})
			},
		},
		seed.Seed{
			Name: "CreateCollectionStatuses",
			Run: func(db *gorm.DB) error {
				return CreateCollectionStatuses(db, []models.CollectionStatus{
					models.NewCollectionStatus(1, "Owned"),
					models.NewCollectionStatus(2, "SaleOrTrade"),
					models.NewCollectionStatus(3, "Trade"),
					models.NewCollectionStatus(4, "Sale"),
				})
			},
		},
	}
}
