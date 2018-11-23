package oauth

import (
	"net/http"

	"github.com/spring2go/gravitee/models"
	"github.com/spring2go/gravitee/oauth/tokentypes"
)

func (s *Service) clientCredentialsGrant(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error) {
	// Get the scope string
	scope, err := s.GetScope(r.Form.Get("scope"))
	if err != nil {
		return nil, err
	}

	// Create a new access token
	accessToken, err := s.GrantAccessToken(
		client,
		nil,                             // empty user
		s.cfg.Oauth.AccessTokenLifetime, // expires in
		scope,
	)
	if err != nil {
		return nil, err
	}

	// Create response
	accessTokenResponse, err := NewAccessTokenResponse(
		accessToken,
		nil, // refresh token
		s.cfg.Oauth.AccessTokenLifetime,
		tokentypes.Bearer,
	)
	if err != nil {
		return nil, err
	}

	return accessTokenResponse, nil
}
