package songdao

import (
	"app/database"
	"app/structs"
)

// CreateSongForBand creates a song for a band in the database itself
func CreateSongForBand(song *structs.Song) (int, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(song).Insert()
	return song.ID, err
}
