package structs

// SongInstrument is used to link multiple instruments to a single song, and give it a PDF
type SongInstrument struct {
	tableName struct{} `pg:"songinstruments"`

	SongID       int    `pg:"song_id"`
	InstrumentID int    `pg:"instrument_id"`
	PDFLocation  string `pg:"pdf_location"`
}
