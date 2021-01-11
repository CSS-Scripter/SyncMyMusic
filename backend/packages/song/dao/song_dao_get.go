package songdao

import (
	"app/database"
	"app/structs"
)

// GetSongsForBand gets all songs for a band
func GetSongsForBand(bandID int) ([]structs.Song, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := []structs.Song{}
	err := db.Model(&result).Where("band_id = ?", bandID).Select()

	return result, err
}

// GetSong gets a single song by given ID
func GetSong(songID int) (structs.Song, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := structs.Song{}
	err := db.Model(&result).Where("id = ?", songID).Select()

	return result, err
}

// GetInstrumentsForSong gets all instruments that have a part within the song.
func GetInstrumentsForSong(songID int) ([]structs.SongInstrument, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := []structs.SongInstrument{}
	err := db.Model(&result).Where("song_id = ?", songID).Select()

	return result, err
}

// GetBandIDFromSongID returns the bandID of the given song
func GetBandIDFromSongID(songID int) (int, error) {
	db := database.ConnectToDatabase()
	defer db.Close()

	result := &structs.Song{ID: songID}
	err := db.Model(result).WherePK().Select()
	return result.BandID, err
}
