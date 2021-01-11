package rehearsal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRehearsalManager(t *testing.T) {
	var expected int = 0

	var actual int = len(rehearsals)

	require.Equal(t, expected, actual, "List should be empty")
}

func TestStartRehearsal_RehearsalsHasLengthOne(t *testing.T) {
	var id int = 123
	var userID int = 321
	var expectedRehearsalsLength int = 1

	var rehearsal = startRehearsal(id, userID)

	var actualRehearsalsLength = len(rehearsals)

	require.NotNil(t, rehearsal, "Rehearsal not created")
	require.Equal(t, expectedRehearsalsLength, actualRehearsalsLength, "Length of rehearsals is not equal to 1")
}

func TestGetRehearsal_ReturnRehearsal(t *testing.T) {
	var id int = 123

	var actual, ok = getRehearsal(id)

	require.NotNil(t, actual, "Rehearsal does not exist")
	require.True(t, ok, "Rehearsal does not exist")
}

func TestDeleteRehearsal_GetRehearsalShouldReturnNil(t *testing.T) {
	var id int = 123

	DeleteRehearsal(id)

	var _, ok = getRehearsal(id)
	require.False(t, ok, "Rehearsal not succesfully deleted")
}

func TestNewPool_ShouldReturnANewPool(t *testing.T) {
	var rehearsalID int = 123
	var ownerID int = 321

	var actual = NewPool(rehearsalID, ownerID)

	require.NotNil(t, actual, "Rehearsal not created")
}
