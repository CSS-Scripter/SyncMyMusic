package structs

// Band structure defines the data stored in a band
type Band struct {
	tableName struct{} `pg:"bands"`

	ID      int    `pg:"id,pk"`
	Name    string `json:"name" pg:"band_name"`
	OwnerID int    `pg:"owner_id"`
}
