package session

import (
	"encoding/gob"
	"errors"
	"net/http"

	// "github.com/gorilla/sessions"
	"github.com/gorilla/sessions"
	"github.com/spring2go/gravitee/config"
)

// Service wraps session functionality
type Service struct {
	sessionStore   sessions.Store
	sessionOptions *sessions.Options
	session        *sessions.Session
	r              *http.Request
	w              http.ResponseWriter
}

// UserSession has user data stored in a session after logging in
type UserSession struct {
	ClientID     string
	Username     string
	AccessToken  string
	RefreshToken string
}

var (
	// StorageSessionName ...
	StorageSessionName = "gravitee_oauth2_server_session"
	// UserSessionKey ...
	UserSessionKey = "gravitee_oauth2_server_user"
	// ErrSessionNotStarted ...
	ErrSessionNotStarted = errors.New("Session not started")
)

func init() {
	// Register a new datatype for storage in sessions
	gob.Register(new(UserSession))
}

// NewService returns a new Service instance
func NewService(config *config.Config, sessionStore sessions.Store) *Service {
	return &Service{
		// Session cookie storage
		sessionStore: sessionStore,
		// Session options
		sessionOptions: &sessions.Options{
			Path:     config.Session.Path,
			MaxAge:   config.Session.MaxAge,
			HttpOnly: config.Session.HTTPOnly,
		},
	}
}

// SetSessionService sets the request and responseWriter on the session service
func (s *Service) SetSessionService(r *http.Request, w http.ResponseWriter) {
	s.r = r
	s.w = w
}

// StartSession starts a new session. This method must be called before other
// public methods of this struct as it sets the internal session object
func (s *Service) StartSession() error {
	session, err := s.sessionStore.Get(s.r, StorageSessionName)
	if err != nil {
		return err
	}
	s.session = session
	return nil
}

// GetUserSession returns the user session
func (s *Service) GetUserSession() (*UserSession, error) {
	// Make suer StartSession has been called
	if s.session == nil {
		return nil, ErrSessionNotStarted
	}

	// Retrieve our user session struct and type-assert it
	userSession, ok := s.session.Values[UserSessionKey].(*UserSession)
	if !ok {
		return nil, errors.New("User session type assertion error")
	}

	return userSession, nil
}

// SetUserSession saves the user session
func (s *Service) SetUserSession(userSession *UserSession) error {
	// Make sure StartSession has been called
	if s.session == nil {
		return ErrSessionNotStarted
	}

	// Set a new user session
	s.session.Values[UserSessionKey] = userSession
	return s.session.Save(s.r, s.w)
}

// ClearUserSession deletes the user session
func (s *Service) ClearUserSession() error {
	// Make sure StartSession has been called
	if s.session == nil {
		return ErrSessionNotStarted
	}

	// Delete the user session
	delete(s.session.Values, UserSessionKey)
	return s.session.Save(s.r, s.w)
}

// SetFlashMessage sets a flash message,
// useful for displaying an error after 302 redirection
func (s *Service) SetFlashMessage(msg string) error {
	// Make sure StartSession has been called
	if s.session == nil {
		return ErrSessionNotStarted
	}

	// Add the flash message
	s.session.AddFlash(msg)
	return s.session.Save(s.r, s.w)
}

// GetFlashMessage returns the first flash message
func (s *Service) GetFlashMessage() (interface{}, error) {
	// Make sure StartSession has been called
	if s.session == nil {
		return nil, ErrSessionNotStarted
	}

	// Get the last flash message from the stack
	if flashes := s.session.Flashes(); len(flashes) > 0 {
		// We need to save the session, otherwise the flash message won't be removed
		s.session.Save(s.r, s.w)
		return flashes[0], nil
	}

	// No flash messages in the stack
	return nil, nil
}

// Close stops any running services
func (s *Service) Close() {}
