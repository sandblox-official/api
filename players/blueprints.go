package players

import "gorm.io/gorm"

//Players is the sql blueprint for a stored player
type Players struct {
	gorm.Model
	ID       int
	Level    float64
	Data     string //Will return a javascript object
	Username string
	HashPass string
}
