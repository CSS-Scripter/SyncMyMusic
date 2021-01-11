package structs

// Song struct defines the values given to a song
type Song struct {
	tableName struct{} `pg:"songs"`

	ID     int    `pg:"id,pk"`
	BandID int    `pg:"band_id"`
	Name   string `json:"name" pg:"song_name"`
	Artist string `json:"artist" pg:"artist_name"`
}
