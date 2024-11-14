package entities

type PlantCondition struct {
	ID               int
	PlantID          int
	Plant            Plant
	Date             string
	MoistureLevel    float32
	SunlightExposure string
	Temperature      float32
	Notes            string
}
