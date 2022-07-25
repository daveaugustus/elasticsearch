package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateARandomUser(t *testing.T) {
	emptyUser := &User{}
	require.Empty(t, emptyUser.FullName)
	require.Empty(t, emptyUser.Email)
	require.NotNil(t, emptyUser.Created)
	require.NotNil(t, emptyUser.DOB)
	require.NotNil(t, emptyUser.Updated)

	user := emptyUser.CreateARandomUser()
	require.NotEmpty(t, user.FullName)
	require.NotEmpty(t, user.Email)
	require.NotNil(t, user.Created)
	require.NotNil(t, user.DOB)
	require.NotNil(t, user.Updated)
}

func TestCreateRandomUsers(t *testing.T) {
	users := CreateRandomUsers(5)
	require.Equal(t, len(users), 5)
	require.NotEmpty(t, users)
	require.NotEmpty(t, users[0].FullName)
	require.NotEmpty(t, users[4].Email)
}
