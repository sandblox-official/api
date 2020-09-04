package worlds

import "gorm.io/gorm"

//Worlds is the sql blueprint for a stored world
type Worlds struct {
	gorm.Model
	Name  string
	Owner string
	//InstructionsForBuild will be an odd way to
	//store world data, but can return direct js
	//to simplify things for now
	InstructionsForBuild string
}
