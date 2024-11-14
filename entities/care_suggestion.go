package entities

type CareSuggestion struct {
	ID         int
	PlantID    int
	Plant      Plant
	Suggestion string
}
