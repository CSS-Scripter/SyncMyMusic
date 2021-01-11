package songdao

import (
	"app/database"
	"app/errors"
	"app/structs"
)

// UpdateSong updates the song in the database
func UpdateSong(song *structs.Song) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	result, err := db.Model(song).WherePK().Update()

	if result.RowsAffected() == 0 {
		return new(errors.NoDatabaseEntryError)
	}

	return err
}
