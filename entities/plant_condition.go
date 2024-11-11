package entities

type PlantCondition struct {
	ID               int
	PlantID          int
	Date             string
	MoistureLevel    float32
	SunlightExposure string
	Temperature      float32
	Notes            string
}
