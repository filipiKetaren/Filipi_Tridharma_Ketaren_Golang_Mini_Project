package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`

	Plants []Plant `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Plant struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	PlantName string `gorm:"not null"`
	Species   string `gorm:"not null"`
	Location  string `gorm:"not null"`

	Conditions      []PlantCondition `gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	CareSuggestions []CareSuggestion `gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
}

type PlantCondition struct {
	ID               int    `gorm:"primaryKey"`
	PlantID          int    `gorm:"not null"`
	Plant            Plant  `gorm:"foreignKey:PlantID;references:ID;constraint:OnDelete:CASCADE"`
	Date             string `gorm:"not null"`
	MoistureLevel    float32
	SunlightExposure string
	Temperature      float32
	Notes            string
}

type CareSuggestion struct {
	ID         int    `gorm:"primaryKey"`
	PlantID    int    `gorm:"not null"`
	Suggestion string `gorm:"type:text;not null"`
	Plant      Plant  `gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE;"`
}
