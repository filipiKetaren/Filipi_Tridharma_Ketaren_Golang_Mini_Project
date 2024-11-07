package entities

import "time"

type PlantCondition struct {
	ID               int
	PlantID          int
	Date             time.Time
	MoistureLevel    float32
	SunlightExposure string
	Temperature      float32
	Notes            string
}
