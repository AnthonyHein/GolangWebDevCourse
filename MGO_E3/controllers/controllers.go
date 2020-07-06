package controllers

import (
    "GolangWebDevCourse/MGO_E3/sessions"
    "GolangWebDevCourse/MGO_E3/models"
    "golang.org/x/crypto/bcrypt"
    "github.com/satori/go.uuid"
    "html/template"
    "net/http"
    "time"
)

type Controller struct {
    Session *sessions.SessionMeta
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func NewController() *Controller {
    return &Controller{
        Session : sessions.NewSession(),
    }
}

func (con Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := con.Session.GetUser(w, req)
	con.Session.ShowSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (con Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := con.Session.GetUser(w, req)
	if !con.Session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	con.Session.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func (con Controller) Signup(w http.ResponseWriter, req *http.Request) {
	if con.Session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := con.Session.DBUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessions.SessionLength
		http.SetCookie(w, c)
		con.Session.DBSessions[c.Value] = models.Session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		con.Session.DBUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	con.Session.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (con Controller) Login(w http.ResponseWriter, req *http.Request) {
	if con.Session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := con.Session.DBUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessions.SessionLength
		http.SetCookie(w, c)
		con.Session.DBSessions[c.Value] = models.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	con.Session.ShowSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (con Controller) Logout(w http.ResponseWriter, req *http.Request) {
	if !con.Session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(con.Session.DBSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(con.Session.DBSessionsCleaned) > (time.Second * 30) {
		go con.Session.CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
