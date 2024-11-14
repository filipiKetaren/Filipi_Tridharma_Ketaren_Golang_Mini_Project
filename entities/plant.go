package entities

type Plant struct {
	ID        int
	UserID    int
	User      User
	PlantName string
	Species   string
	Location  string
}

type PlantWithUser struct {
	Plant struct {
		ID        int
		UserID    int
		User      User
		PlantName string
		Species   string
		Location  string
	}
}
