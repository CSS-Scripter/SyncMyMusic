package structs

// Instrument structure defines the values for an instrument
type Instrument struct {
	tableName struct{} `pg:"instruments"`

	ID     int    `pg:"id,pk"`
	BandID int    `pg:"band_id"`
	Name   string `pg:"instrument_name"`
}
