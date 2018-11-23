package web

import (
	"net/http"

	"github.com/spring2go/gravitee/config"
	"github.com/spring2go/gravitee/oauth"
	"github.com/spring2go/gravitee/session"
)

// Service struct keeps variables for reuse
type Service struct {
	cfg            *config.Config
	oauthService   oauth.ServiceInterface
	sessionService session.ServiceInterface
}

// NewService returns a new Service instance
func NewService(cfg *config.Config, oauthService oauth.ServiceInterface, sessionService session.ServiceInterface) *Service {
	return &Service{
		cfg:            cfg,
		oauthService:   oauthService,
		sessionService: sessionService,
	}
}

// GetConfig returns config.Config instance
func (s *Service) GetConfig() *config.Config {
	return s.cfg
}

// GetOauthService returns oauth.Service instance
func (s *Service) GetOauthService() oauth.ServiceInterface {
	return s.oauthService
}

// GetSessionService returns session.Service instance
func (s *Service) GetSessionService() session.ServiceInterface {
	return s.sessionService
}

// Close stops any running services
func (s *Service) Close() {}

func (s *Service) setSessionService(r *http.Request, w http.ResponseWriter) {
	s.sessionService.SetSessionService(r, w)
}
