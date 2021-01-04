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

// Condition of the items in collection
type Condition struct {
	ID   int
	Name string
}

// Case type of a collection item
type Case struct {
	ID   int
	Name string
}

// Status of a collection item
type Status struct {
	ID   int
	Name string
}
