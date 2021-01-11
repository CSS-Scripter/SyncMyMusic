package musician

import (
	dao "app/packages/musician/dao"
	"app/structs"

	"golang.org/x/crypto/bcrypt"
)

// CreateMusician is used to create a Musician (User) for the application.
func CreateMusician(musician *structs.Musician) error {
	musician.Password, _ = hashPassword(musician.Password)
	_, err := dao.CreateMusician(musician)
	return err
}

// UpdateMusician is used to change Musician details.
func UpdateMusician(musician structs.Musician, id int) error {
	musician.Password, _ = hashPassword(musician.Password)
	return dao.UpdateMusician(musician, id)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// AuthenticatUser checks if provided password matches known password
func AuthenticatUser(email, password string) bool {
	musician, err := dao.GetMusicianByEmail(email)
	if err != nil {
		return false
	}
	return checkPasswordHash(password, musician.Password)
}

// GetMusicianByEmail returns a musician by its email
func GetMusicianByEmail(email string) (structs.Musician, error) {
	return dao.GetMusicianByEmail(email)
}

// GetMusicianByID returns a musician by its ID
func GetMusicianByID(musicianID int) (structs.Musician, error) {
	return dao.GetMusicianByID(musicianID)
}
