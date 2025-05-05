package models

type User struct {
	ID        uint    `json:"id"`         // PRIMARY KEY, SERIAL/BIGSERIAL
	FirstName string  `json:"first_name"` // VARCHAR(100) NOT NULL
	LastName  string  `json:"last_name"`  // VARCHAR(100) NOT NULL
	Email     string  `json:"email"`      // VARCHAR(100) UNIQUE NOT NULL
	Password  string  `json:"password"`   // VARCHAR(100) NOT NULL, MIN 6 символов (валидация в коде)
	City      string  `json:"city"`       // VARCHAR(100), необязательное поле
	Plants    []Plant `json:"plants"`     // Связь "один-ко-многим" (user → plants), реализуется вручную
}
