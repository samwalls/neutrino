// Current sessions and users

package main

import (
    "math/rand"
    "errors"
)

type Session struct {
	users []User 
	fileName string
	id string
}

type User struct {
	username string
}

type CursorPos struct {
	line string
	column string
}

const (
	idLength int  = 10;
)

var sessions= make([]Session, 0)

func newSessionId() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
    b := make([]rune, idLength)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// Creates a new session
func NewSession(createrUsername string, fileName string, initialFileData string) Session  {
	cUser := User{username: createrUsername}
	session := Session{fileName: fileName, id: newSessionId(), users: make([]User, 1)}
	session.users[0] =  cUser
	return session
}

func GetSessionById(id string) (s Session, err error) {
	for _, session := range sessions {
		if session.id == id {
			return session, nil
		}
	}

	return s, errors.New("No session with id found")
}

func GetUsernamesForSession(sessionId string) (userIds []string, err error) {
	session, err := GetSessionById(sessionId)
	if err == nil {
		return userIds, err
	}

	users := session.users
	ids := make([]string, 0)

	for idx, user := range users {
		ids[idx] = user.username
	}

	return ids, nil
}

func main() {
}
