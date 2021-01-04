package seeds

import (
	"github.com/zorrokid/movieCollection/pkg/seed"
	"gorm.io/gorm"
)

// All returns instances of seeds
func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateConditions",
			Run: func(db *gorm.DB) error {
				return CreateCondition(db, []string{
					"Poor - Slightly Damaged",
					"Bad - Damaged",
					"Good",
					"Fair",
					"Excellent",
					"Still Wrapped",
				})
			},
		},
	}
}
