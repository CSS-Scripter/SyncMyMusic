package structs

// BandMember is the struct that connects musicians to bands
type BandMember struct {
	tableName struct{} `pg:"bandmembers,alias:bandmember"`

	MusicianID   int         `pg:"musician_id"`
	Musician     *Musician   `pg:"rel:has-one"`
	BandID       int         `pg:"band_id"`
	Band         *Band       `pg:"rel:has-one"`
	InstrumentID int         `pg:"instrument_id"`
	Instrument   *Instrument `pg:"rel:has-one"`
}
