package session

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/dayaftereh/discover/server/utils"
	"github.com/gorilla/sessions"
)

const (
	SessionStorageKey = "http:session"
)

type Manager struct {
	sessionName string
	store       *sessions.CookieStore
}

func NewSessionManager() (*Manager, error) {
	bytes, err := utils.RandBytes(256) // 256 random bytes
	if err != nil {
		return nil, errors.Wrap(err, "fail to generated random bytes for cookie store")
	}

	// register the http session
	gob.Register(&HttpSession{})

	manager := &Manager{
		sessionName: "discover",
		store:       sessions.NewCookieStore(bytes),
	}

	return manager, nil
}

func (manager *Manager) Get(response http.ResponseWriter, request *http.Request) (*HttpSession, error) {
	// get the session from store
	session, err := manager.store.Get(request, manager.sessionName)
	if err != nil {
		return nil, err
	}

	// check if a new session
	if session.IsNew {
		id, err := utils.RandString(64)
		if err != nil {
			return nil, err
		}

		log.Println("creating new http-session with id [", id, "]")

		// create a new session
		httpSession := &HttpSession{
			Id: id,
		}

		// store the new created session
		session.Values[SessionStorageKey] = httpSession
		// save the session
		err = session.Save(request, response)
		if err != nil {
			return nil, err
		}
	}

	// get the http-session from values
	httpSession, ok := session.Values[SessionStorageKey].(*HttpSession)
	if !ok {
		return nil, errors.Errorf("unable to find value [ %s ] in session", SessionStorageKey)
	}

	return httpSession, nil
}
