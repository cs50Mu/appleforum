package views

import (
	"appleforum/models/db"
	"appleforum/templates"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// IndexHandler handles index page
var IndexHandler = SessionHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	//
	userID := session.UserID
	var userName string
	err := db.QueryRow("SELECT name FROM users WHERE id = ?", userID).Scan(&userName)
	if err != nil {
		fmt.Fprint(w, "Hello, stranger!")
		return
	}
	fmt.Fprintf(w, "Hello, %v", userName)
})

// LoginHandler handles login req
var LoginHandler = SessionHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	// get redirect_url
	redirectURL := r.URL.Query().Get("next")
	if redirectURL == "" {
		redirectURL = "/"
	}
	log.Printf("redirect url at first: %v", redirectURL)
	// authenticate user
	// if success, set cookie && set session && redirect to redirect_url
	if r.Method == http.MethodPost {
		// check username && passwd
		userName := r.PostFormValue("user_name")
		passwd := r.PostFormValue("passwd")
		if err := session.Authenticate(userName, passwd); err != nil {
			http.Redirect(w, r, "/login?next="+url.QueryEscape(redirectURL), http.StatusSeeOther)
			return
		}
		log.Printf("success redirect url: %v", redirectURL)
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		return
	}
	// if not, render login form
	if session.UserID != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	templates.Render(w, "login", nil)
})

// LogoutHandler handles logout req
var LogoutHandler = SessionHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	// remove session
	sessionID := session.SessionID
	db.Exec("DELETE FROM sessions WHERE session_id = ?", sessionID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
})

// Session is a http session
type Session struct {
	SessionID string
	UserID    int64
}

// Authenticate a user
func (sess *Session) Authenticate(userName, passwd string) error {
	// check user from db
	var userID int64
	var pass string
	err := db.QueryRow("SELECT id,passwd FROM users where name = ?", userName).Scan(&userID, &pass)
	if err != nil {
		return errors.New("user name or password is not correct")
	}
	log.Printf("pass:%v, passwd:%v", pass, passwd)
	if pass != passwd {
		return errors.New("user name or password is not correct")
	}
	// if success, update session
	db.Exec("UPDATE sessions SET user_id = ? WHERE session_id = ?", userID, sess.SessionID)
	return nil
}

// SessionKey is the session key set in cookie
// about naming convention, refer to https://stackoverflow.com/questions/22688906/go-naming-conventions-for-const
const SessionKey = "sessionID"

func getSession(w http.ResponseWriter, r *http.Request) Session {
	// get sessionid from cookie
	var session Session
	cookie, err := r.Cookie(SessionKey)
	if err != nil {
		// cookie not found, create new session
		session = newSession(w, r)
		return session
	}
	// check sessionid from database
	sessionID := cookie.Value
	var userID int64
	err = db.QueryRow("SELECT user_id FROM sessions WHERE session_id = ?", sessionID).Scan(&userID)
	if err == sql.ErrNoRows {
		// no matching session, create new session
		// generate session_id
		// if not, redirect to login view
		session = newSession(w, r)
		return session
	}
	// if success, set cookie && session, continue
	session = Session{
		SessionID: sessionID,
		UserID:    userID,
	}
	return session
}

func newSession(w http.ResponseWriter, r *http.Request) Session {
	//
	sessionID := randStr(15)
	// insert session record
	db.Exec(`INSERT INTO sessions (session_id, user_id)
			VALUES (?, ?)`, sessionID, 0)
	session := Session{
		SessionID: sessionID,
		UserID:    0,
	}
	cookie := http.Cookie{
		Name:  SessionKey,
		Value: sessionID,
	}
	http.SetCookie(w, &cookie)
	return session
}

// LoginRequiredHandler ensures login first action
var LoginRequiredHandler = func(handler func(w http.ResponseWriter, r *http.Request, session *Session)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get session
		session := getSession(w, r)
		// get url path
		urlPath := r.URL.Path
		log.Printf("path: %v", urlPath)
		// login don't need redirect
		if urlPath == "/login" {
			handler(w, r, &session)
			return
		}
		if r.URL.RawQuery != "" {
			urlPath += "?" + r.URL.RawQuery
		}
		// if user not logged in
		// redirect to login
		if session.UserID == 0 {
			http.Redirect(w, r, "/login?next="+url.QueryEscape(urlPath), http.StatusSeeOther)
			return
		}
		handler(w, r, &session)
	}
}

// SessionHandler is the middleware that handles session
// it take a func with signature of `(w http.ResponseWriter, r *http.Request, session Session)`
// returns a func with signature of `(http.ResponseWriter, *http.Request)` which satisfies
// net.http.HandlerFunc
var SessionHandler = func(handler func(w http.ResponseWriter, r *http.Request, session *Session)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get session
		session := getSession(w, r)
		handler(w, r, &session)
	}
}
