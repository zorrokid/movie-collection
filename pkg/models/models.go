package models

// Compilation contains multiple movies
type Compilation struct {
	ID            int
	OriginalTitle string
	LocalTitle    string
	Country       string
	CaseID        int
	Barcode       string
	ConditionID   int
	Notes         string
	StatusID      int
	Checked       bool
	Edition       string
	Discs         int
	HasSlipCover  bool
	HasHologram   bool
	IsRental      bool
}

type nonAutoIncrementModel struct {
	ID uint `gorm:"primary_key;autoincrement:false"`
}

// Condition of the items in collection
type Condition struct {
	ID   uint `gorm:"primary_key;autoincrement:false"`
	Name string
}

// NewCondition constructs Condition object
func NewCondition(id uint, name string) Condition {
	return Condition{
		ID:   id,
		Name: name,
	}
}

// CaseType type of a collection item
type CaseType struct {
	ID   uint `gorm:"primary_key;autoincrement:false"`
	Name string
}

// NewCaseType constructs Case object
func NewCaseType(id uint, name string) CaseType {
	return CaseType{
		ID:   id,
		Name: name,
	}
}

// CollectionStatus of a collection item
type CollectionStatus struct {
	ID   uint `gorm:"primary_key;autoincrement:false"`
	Name string
}

// NewCollectionStatus constructs Status object
func NewCollectionStatus(id uint, name string) CollectionStatus {
	return CollectionStatus{
		ID:   id,
		Name: name,
	}
}
