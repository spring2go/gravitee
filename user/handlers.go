package user

import (
	"net/http"

	"github.com/spring2go/gravitee/oauth/roles"
	"github.com/spring2go/gravitee/util/response"
)

// Handles create user requests (GET /v1/health)
func (s *Service) createUser(w http.ResponseWriter, r *http.Request) {

	// Parse the form so r.Form becomes available
	if err := r.ParseForm(); err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// Check username or pasword empty
	if username == "" || password == "" {
		response.Error(w, "username or password can't be empty", http.StatusBadRequest)
		return
	}

	// Check user existence
	if s.oauthService.UserExists(username) {
		response.Error(w, "username taken", http.StatusBadRequest)
		return
	}

	// Create a user
	_, err := s.oauthService.CreateUser(
		roles.User, // role ID
		username,   // username
		password,   // password
	)

	if err != nil {
		response.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.WriteJSON(w, map[string]interface{}{
		"success": true,
	}, 200)
}
