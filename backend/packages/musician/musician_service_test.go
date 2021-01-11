package musician

import (
	"app/database"
	"app/structs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthenticateUser_ShouldReturnTrue(t *testing.T) {
	password := "geheim"
	musician := &structs.Musician{
		Email:    "test@email.nl",
		Password: password,
		Name:     "test",
	}
	CreateMusician(musician)
	expected := true

	actual := AuthenticatUser(musician.Email, password)
	db := database.ConnectToDatabase()
	defer db.Close()
	db.Model(musician).WherePK().Delete()

	require.Equal(t, expected, actual)
}

func TestAuthenticateUser_ShouldReturnFalse(t *testing.T) {
	rightPassword := "geheim"
	wrongPassword := "bekend"
	musician := &structs.Musician{
		Email:    "test@email.nl",
		Password: rightPassword,
		Name:     "test",
	}
	CreateMusician(musician)
	expected := false

	actual := AuthenticatUser(musician.Email, wrongPassword)
	db := database.ConnectToDatabase()
	defer db.Close()
	db.Model(musician).WherePK().Delete()

	require.Equal(t, expected, actual)
}
