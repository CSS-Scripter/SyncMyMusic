package songdao

import (
	"app/database"
	"app/structs"
)

// DeleteSong deletes an song from a band's repertoire
func DeleteSong(songID int) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(&structs.SongInstrument{}).Where("song_id = ?", songID).Delete()
	if err != nil {
		return err
	}

	_, err = db.Model(&structs.Song{}).Where("id = ?", songID).Delete()
	return err
}
