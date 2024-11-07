package entities

type Reminder struct {
	ID           uint
	PlantID      uint
	Plant        Plant
	ReminderText string
	IsRecurring  bool
	IntervalDays int
}
