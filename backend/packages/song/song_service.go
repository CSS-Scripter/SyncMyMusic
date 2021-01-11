package song

import (
	dao "app/packages/song/dao"
	"app/structs"
	"fmt"
	"os"
)

// CreateSongForBand creates an song for a band
func CreateSongForBand(bandID int, song *structs.Song) (int, error) {
	song.BandID = bandID
	return dao.CreateSongForBand(song)
}

// GetSongsForBand gets all songs for a certain band
func GetSongsForBand(bandID int) ([]structs.Song, error) {
	return dao.GetSongsForBand(bandID)
}

// DeleteSong deletes an song
func DeleteSong(songID int) error {
	bandID, err := dao.GetBandIDFromSongID(songID)
	if err != nil {
		return err
	}
	err = dao.DeleteSong(songID)
	if err != nil {
		return err
	}
	err = os.RemoveAll(fmt.Sprintf("files/%v/%v", bandID, songID))
	return err
}

// UpdateSong updates a song in the database
func UpdateSong(song *structs.Song) error {
	return dao.UpdateSong(song)
}

// GetInstrumentsForSong gets all instruments that have a part inside a song.
func GetInstrumentsForSong(songID int) ([]structs.SongInstrument, error) {
	return dao.GetInstrumentsForSong(songID)
}

// GetBandIDFromSong returns the bandID of the given song
func GetBandIDFromSong(songID int) (int, error) {
	return dao.GetBandIDFromSongID(songID)
}
