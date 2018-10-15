package session

import (
	"../../utils"
	"fmt"
	"github.com/gorilla/sessions"
	"log"
)

const (
	SESSION_STORE_KEY = "http:session"
)

type SessionManager struct {
	sessionName string
	store       *sessions.CookieStore
}

func NewSessionManager() (*SessionManager, error) {
	sessionManager := &SessionManager{}
	err := sessionManager.init()
	if err != nil {
		return nil, err
	}
	return sessionManager, nil
}

func (sessionManager *SessionManager) init() error {
	bytes, err := utils.RandBytes()
	if err != nil {
		return err
	}
	sessionManager.store = sessions.NewCookieStore(bytes)
}

func (sessionManager *SessionManager) Get(response http.ResponseWriter, request *http.Request) (*HttpSession, error) {
	// get the session from store
	session, err := sessionManager.store.Get(r, sessionManager.sessionName)
	if err != nil {
		return nil, err
	}

	// check if a new session
	if session.IsNew {
		id, err := utils.RandString(64)
		if err != nil {
			return err
		}

		log.Println("creating new http-session with id [", id, "]")

		// create a new session
		httpSession := &HttpSession{
			Id: id,
		}

		// store the new created session
		session.Values[SESSION_STORE_KEY] = httpSession
		// save the session
		err = session.Save(r, w)
		if err != nil {
			return "", err
		}
	}

	// get the http-session from values
	httpSession, ok := session.Values[SESSION_STORE_KEY].(*HttpSession)
	if !ok {
		return "", fmt.Errorf("unable to find value [ %s ] in session", SESSION_STORE_KEY)
	}

	return httpSession, nil
}
