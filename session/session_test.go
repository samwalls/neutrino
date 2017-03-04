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

func TestSetCursorInfo(t *testing.T) {
	session := NewSession("james", "test.txt", "pass")
	pos := FilePos{line: 10, column: 20}
	sel := FileSelection{
		start: FilePos{line: 1, column: 2},
		end: FilePos{line: 3, column: 4},
	}

	err := SetCursorPosAndSelection(session.id, "james", pos, sel)
	assert.Equal(t, err, nil, "")
	session, _ = GetSessionById(session.id)
	user := session.users[0]
	assert.Equal(t, user.cursorPos.line, 10, "")
	assert.Equal(t, user.cursorPos.column, 20, "")

	assert.Equal(t, user.cursorSelection.start.line, 1, "")
	assert.Equal(t, user.cursorSelection.start.column, 2, "")

	assert.Equal(t, user.cursorSelection.end.line, 3, "")
	assert.Equal(t, user.cursorSelection.end.column, 4, "")
}