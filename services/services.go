package services

import (
	"reflect"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/spring2go/gravitee/config"
	"github.com/spring2go/gravitee/health"
	"github.com/spring2go/gravitee/oauth"
	"github.com/spring2go/gravitee/session"
	"github.com/spring2go/gravitee/user"
	"github.com/spring2go/gravitee/web"
)

func init() {

}

var (
	// HealthService ...
	HealthService health.ServiceInterface

	// OauthService ...
	OauthService oauth.ServiceInterface

	// UserService ...
	UserService user.ServiceInterface

	// WebService ...
	WebService web.ServiceInterface

	// SessionService ...
	SessionService session.ServiceInterface
)

// UseHealthService sets the health service
func UseHealthService(h health.ServiceInterface) {
	HealthService = h
}

// UseOauthService sets the oAuth service
func UseOauthService(o oauth.ServiceInterface) {
	OauthService = o
}

// UseUserService sets the user service
func UseUserService(u user.ServiceInterface) {
	UserService = u
}

// UseWebService sets the web service
func UseWebService(w web.ServiceInterface) {
	WebService = w
}

// UseSessionService sets the session service
func UseSessionService(s session.ServiceInterface) {
	SessionService = s
}

// Init starts up all services
func Init(cfg *config.Config, db *gorm.DB) error {
	if nil == reflect.TypeOf(HealthService) {
		HealthService = health.NewService(db)
	}

	if nil == reflect.TypeOf(OauthService) {
		OauthService = oauth.NewService(cfg, db)
	}

	if nil == reflect.TypeOf(UserService) {
		UserService = user.NewService(OauthService)
	}

	if nil == reflect.TypeOf(SessionService) {
		// note: default session store is CookieStore
		SessionService = session.NewService(cfg, sessions.NewCookieStore([]byte(cfg.Session.Secret)))
	}

	if nil == reflect.TypeOf(WebService) {
		WebService = web.NewService(cfg, OauthService, SessionService)
	}

	return nil
}

// Close closes any open services
func Close() {
	HealthService.Close()
	UserService.Close()
	OauthService.Close()
	WebService.Close()
	SessionService.Close()
}
