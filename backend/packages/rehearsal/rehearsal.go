package rehearsal

// Rehearsal Struct for a rehearsal
type Rehearsal struct {
	ID   int
	pool *Pool
}

// NewRehearsal consturctor for rehearsal
func NewRehearsal(bandID int, OwnerID int) *Rehearsal {
	return &Rehearsal{
		ID:   bandID,
		pool: NewPool(bandID, OwnerID),
	}
}
