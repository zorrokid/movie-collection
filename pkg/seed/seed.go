package seed

import "gorm.io/gorm"

// Seed for running the database seeds. Based on https://ieftimov.com/post/simple-golang-database-seeding-abstraction-gorm/
type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}
