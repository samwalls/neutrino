// Current sessions and users

package session

import (
    "math/rand"
    "errors"
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

type Session struct {
	Users []User 
	FileName string
	Id string
	HashedPassword []byte
}

type User struct {
	Username string
	CursorPos FilePos
	CursorSelection FileSelection
}

type FilePos struct {
	Line int
	Column int
}

type FileSelection struct {
	Start FilePos
	End FilePos
}

const (
	idLength int  = 10;
)

var sessions= make([]Session, 0)

func newSessionId() string {
	fmt.Println()

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
    b := make([]rune, idLength)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func toByteArray(str string) []byte {
	return []byte(str)
}

func ValidPassword(sessionId string, plainPassword string) bool {
	session, err := GetSessionById(sessionId)
	if err != nil {
		return false
	}

	if session.HashedPassword == nil {
		return true // No password set for session (passwords are optional)
	}

	err = bcrypt.CompareHashAndPassword(session.HashedPassword, toByteArray(plainPassword))
	return err == nil // No error => Valid password
}

// Creates a new session
func NewSession(createrUsername string, fileName string, plainPassword string) Session  {
	cUser := User{Username: createrUsername}
	hashedPassword, _ := bcrypt.GenerateFromPassword(toByteArray(plainPassword), 5) 
	if len(plainPassword) == 0 {
		hashedPassword = nil
	}
	session := Session{FileName: fileName, Id: newSessionId(),Users: make([]User, 1), HashedPassword: hashedPassword}
	session.Users[0] =  cUser

	// Add session to list of all sessions (in memory)
	sessions = append(sessions, session)
	return session
}

func AddUserToSession(sessionId string, username string) (u User, e error) {
	user := User{Username: username}
	session, err := GetSessionById(sessionId)
	if err != nil {
		return u, errors.New("Failed to find session with id")
	}

	session.Users = append(session.Users, user)
	setSessionAtId(sessionId, session)
	return user, nil
}

func setSessionAtId(sessionId string, session Session) {
	for idx, sess := range sessions {
		if sess.Id == sessionId {
			sessions[idx] = session
			return 
		}
	}
}

func GetSessionById(id string) (s Session, err error) {
	for _, session := range sessions {
		if session.Id == id {
			return session, nil
		} 
	}   

	return s, errors.New("No session with id found")
} 

func GetUsernamesForSession(sessionId string) (userIds []string, err error) {
	session, err := GetSessionById(sessionId)
	if err != nil {
		return userIds, err
	}

	users := session.Users
	ids := make([]string, len(users))
 
	for idx, user := range users {
		ids[idx] = user.Username
	}

	return ids, nil
}

func SetCursorPosAndSelection(sessionId string, username string, newCursorPos FilePos, newCursorSelection FileSelection) error {
	session, err := GetSessionById(sessionId)
	if err != nil {
		return err
	}

	users := session.Users

	for idx, user := range users {
		if user.Username == username {
			users[idx].CursorPos = newCursorPos
			users[idx].CursorSelection = newCursorSelection
			session.Users = users 
			setSessionAtId(sessionId, session)
			return nil;
		}
	}

	return errors.New("No user found with that username in the session")
}