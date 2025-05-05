package models

type Plant struct {
	ID          uint   `json:"id"`          // PRIMARY KEY, SERIAL/BIGSERIAL
	Name        string `json:"name"`        // VARCHAR(100) NOT NULL
	Description string `json:"description"` // TEXT
	Watering    string `json:"watering"`    // VARCHAR(50)
	Repotting   string `json:"repotting"`   // VARCHAR(50)
	Breeding    string `json:"breeding"`    // VARCHAR(50)
	UserID      uint   `json:"user_id"`     // NOT NULL, FOREIGN KEY to users(id)
}
