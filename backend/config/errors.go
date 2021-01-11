package config

type errorMessages struct {
	IDIsNotANumber   string
	NoDatabaseEntry  string
	InvalidID        string
	BandDoesNotExist string
}

// ErrorMessages is the object containing all error messages
var ErrorMessages errorMessages

func init() {
	ErrorMessages = errorMessages{
		IDIsNotANumber:   "Given ID was not a number",
		NoDatabaseEntry:  "404: There was no database entry found",
		InvalidID:        "The given ID was not valid",
		BandDoesNotExist: "The given band does not exist",
	}
}
