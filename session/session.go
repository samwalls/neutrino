// Current sessions and users

package main

import (
    "math/rand"
    "fmt"
)

type session struct {
	users []user 
	fileName string
	id string
}

type user struct {
	username string
}

const (
	idLength int  = 10;
)

func newSessionId() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
    b := make([]rune, idLength)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// Creates a new session, returning it's ID
func NewSession(createrUsername string, fileName string, initialFileData string) string {
	cUser := user{username: createrUsername}
	session := session{fileName: fileName, id: newSessionId(), users: make([]user, 1)}
	session.users[0] =  cUser
	return session.id
}

// When a file is changed by a user
func FileChanged(sessionId string, username string, updatedFileContents string) {
	
}

func main() {
}
