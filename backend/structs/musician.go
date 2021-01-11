package structs

type Musician struct {
	tableName struct{} `pg:"musicians,alias:musicians"`

	MusicianID int    `pg:"id"`
	Name       string `json:"username" pg:"musician_name"`
	Email      string `json:"email" pg:"email"`
	Password   string `pg:"password"`
}
