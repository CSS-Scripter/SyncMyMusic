package config

type messages struct {
	BandCreatedMessage string
	BandDeletedMessage string

	SongCreatedMessage string
	SongDeletedMessage string

	InstrumentCreatedMessage string
	InstrumentDeletedMessage string

	MusicianCreatedMessage string
	MusicianUpdatedMessage string

	BandmemberDeletedMessage string
}

// Messages is the object containing all messages
var Messages messages

var createOK = "was created succesfully"
var deleteOK = "was deleted succesfully"

func init() {
	Messages = messages{
		BandCreatedMessage: "Band " + createOK,
		BandDeletedMessage: "Band " + deleteOK,

		SongCreatedMessage: "Song " + createOK,
		SongDeletedMessage: "Song " + deleteOK,

		MusicianCreatedMessage:   "User " + createOK,
		MusicianUpdatedMessage:   "User was updated succesfully",
		InstrumentCreatedMessage: "Instrument " + createOK,
		InstrumentDeletedMessage: "Instrument " + deleteOK,

		BandmemberDeletedMessage: "Bandmember " + deleteOK,
	}
}
