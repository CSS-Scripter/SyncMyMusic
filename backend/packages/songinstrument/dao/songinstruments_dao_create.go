package songinstrumentdao

import (
	"app/database"
	"app/structs"
)

// CreateSongInstrument creates an entry for the songinstrument
func CreateSongInstrument(songInstrument *structs.SongInstrument) error {
	db := database.ConnectToDatabase()
	defer db.Close()

	_, err := db.Model(songInstrument).Insert()
	return err

}
