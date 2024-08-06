package models

import "gorm.io/gorm"

// gotta add the new models class to the db.go

type Tire struct {
	gorm.Model
	TireModelNumber   string
	TireName          string
	TireBrand         string
	TireSize          string
	TirePrice         int64
	TotalTirePrice    int64
	TypeOfTire        string
	TirePly           int
	TireLoadRating_LB int
	TirePressure_PSI  int
	TireTreadDepth    int
}
