package sessions

import (
	"GolangWebDevCourse/MGO_E3/models"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

const SessionLength int = 30

type SessionMeta struct {
	DBUsers map[string]models.User      // user ID, user
    DBSessions map[string]models.Session // session ID, session
    DBSessionsCleaned time.Time
}

func NewSession() *SessionMeta {
    return &SessionMeta{
		DBUsers : make(map[string]models.User),
        DBSessions : make(map[string]models.Session),
        DBSessionsCleaned : time.Now(),
    }
}

func (s SessionMeta) GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	c.MaxAge = SessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if sess, ok := s.DBSessions[c.Value]; ok {
		sess.LastActivity = time.Now()
		s.DBSessions[c.Value] = sess
		u = s.DBUsers[sess.Un]
	}
	return u
}

func (s SessionMeta) AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	sess, ok := s.DBSessions[c.Value]
	if ok {
		sess.LastActivity = time.Now()
		s.DBSessions[c.Value] = sess
	}
	_, ok = s.DBUsers[sess.Un]
	// refresh session
	c.MaxAge = SessionLength
	http.SetCookie(w, c)
	return ok
}

func (s SessionMeta) CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	s.ShowSessions()              // for demonstration purposes
	for k, v := range s.DBSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(s.DBSessions, k)
		}
	}
	s.DBSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	s.ShowSessions()             // for demonstration purposes
}

// for demonstration purposes
func (s SessionMeta) ShowSessions() {
	fmt.Println("********")
	for k, v := range s.DBSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
