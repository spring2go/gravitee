package user

import (
	"github.com/gorilla/mux"
	"github.com/spring2go/gravitee/util/routes"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	// Exported methods
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()
}
