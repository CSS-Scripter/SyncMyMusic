package musiciandao

import (
	"app/database"
	"app/structs"
)

// CreateMusician is used to create a Musician (User) for the application.
func CreateMusician(musician *structs.Musician) (*structs.Musician, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(musician).Insert()
	return musician, err
}

// UpdateMusician is used to change Musician details.
func UpdateMusician(musician structs.Musician, id int) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(&musician).Where("id = ?", id).Update()
	return err
}

// GetMusicianByEmail does exactly what is says
func GetMusicianByEmail(email string) (structs.Musician, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := structs.Musician{Email: email}

	err := db.Model(&result).Where("Email = (?)", email).Last()
	return result, err
}

// GetMusicianByID returns a musician by its ID
func GetMusicianByID(musicianID int) (structs.Musician, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := structs.Musician{MusicianID: musicianID}
	err := db.Model(&result).WherePK().Select()
	result.Password = ""
	return result, err
}

// DeleteMusicianByID deletes an musician by ID
func DeleteMusicianByID(musicianID int) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	model := structs.Musician{MusicianID: musicianID}
	_, err := db.Model(&model).WherePK().Delete()
	return err
}
