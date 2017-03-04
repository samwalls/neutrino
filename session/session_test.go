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
	assert.Equal(t, session.FileName, "test.txt", "")
	users := session.Users
	assert.Equal(t, len(users), 1, "")
	user := users[0]
	assert.Equal(t, user.Username, "james", "")
}

func TestAuthentication(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	assert.True(t, ValidPassword(session.Id, "pass"))
	assert.False(t, ValidPassword(session.Id, "pass2"))
	assert.False(t, ValidPassword(session.Id, ""))
}

func TestGetSessionId(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	_ = NewSession("james", "sdfljhdskfjh", "pass")
	s2, _ := GetSessionById(session.Id)
	assert.Equal(t, s2.FileName, "test.txt", "")
}


func TestGetUserIds(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	sessionId := session.Id

	assert.Equal(t, len(session.Users), 1, "")
	AddUserToSession(sessionId, "john")

	usernames, _ := GetUsernamesForSession(sessionId)

	assert.NotEqual(t, usernames, nil, "")
	assert.Equal(t, len(usernames), 2, "")
	assert.Equal(t, usernames[0], "james", "")
	assert.Equal(t, usernames[1], "john", "")
} 

func TestSetCursorInfo(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	pos := FilePos{Line: 10, Column: 20}
	sel := FileSelection{
		Start: FilePos{Line: 1, Column: 2},
		End: FilePos{Line: 3, Column: 4},
	}

	err := SetCursorPosAndSelection(session.Id, "james", pos, sel)
	assert.Equal(t, err, nil, "")
	session, _ = GetSessionById(session.Id)
	user := session.Users[0]
	assert.Equal(t, user.CursorPos.Line, 10, "")
	assert.Equal(t, user.CursorPos.Column, 20, "")

	assert.Equal(t, user.CursorSelection.Start.Line, 1, "")
	assert.Equal(t, user.CursorSelection.Start.Column, 2, "")

	assert.Equal(t, user.CursorSelection.End.Line, 3, "")
	assert.Equal(t, user.CursorSelection.End.Column, 4, "")
}