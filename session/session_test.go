package session

import (
	"testing"
  	"github.com/stretchr/testify/assert"
  	"fmt"
)

func TestTest(t *testing.T) {
	fmt.Println("")
}

func TestNewSession(t *testing.T) {
	session := NewSession("james", "test.txt", "")
	assert.Equal(t, session.fileName, "test.txt", "")
	users := session.users
	assert.Equal(t, len(users), 1, "")
	user := users[0]
	assert.Equal(t, user.username, "james", "")
}

func TestAuthentication(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	assert.True(t, ValidPassword(session.id, "pass"))
	assert.False(t, ValidPassword(session.id, "pass2"))
	assert.False(t, ValidPassword(session.id, ""))
}

func TestGetSessionId(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	_ = NewSession("james", "sdfljhdskfjh", "pass")
	s2, _ := GetSessionById(session.id)
	assert.Equal(t, s2.fileName, "test.txt", "")
}


func TestGetUserIds(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	sessionId := session.id

	assert.Equal(t, len(session.users), 1, "")
	AddUserToSession(sessionId, "john")

	usernames, _ := GetUsernamesForSession(sessionId)

	assert.NotEqual(t, usernames, nil, "")
	assert.Equal(t, len(usernames), 2, "")
	assert.Equal(t, usernames[0], "james", "")
	assert.Equal(t, usernames[1], "john", "")
} 