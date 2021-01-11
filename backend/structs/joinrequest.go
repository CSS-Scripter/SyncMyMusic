package structs

// JoinRequest represents the request to join a band from an user.
// UserID is the ID of the user wanting to join
// BandID is the ID of the band the user is wanting to join
type JoinRequest struct {
	tableName struct{} `pg:"joinrequests"`

	UserID int `pg:"musician_id" json:"musicianId"`
	BandID int `pg:"band_id" json:"bandId"`
}
