package ft_auth

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type server struct {
	http.Server
	shutdownReq chan bool
}

var token *oauth2.Token

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var setup42 = &oauth2.Config{
	Scopes: []string{"public", "projects", "profile", "tig"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://api.intra.42.fr/oauth/authorize",
		TokenURL: "https://api.intra.42.fr/oauth/token",
	},
	RedirectURL: "http://localhost:25634/callback",
}

func Authenticate(clientId, clientSecret string) bool {
	setup42.ClientID = clientId
	setup42.ClientSecret = clientSecret

	// Get the Token from the json file or the server
	token = retrieveTokenFromFile()
	if token == nil {
		serveToken()
	}

	// Check expiration, validity of the token
	if !token.Valid() {
		return false
	}

	saveTokenOnFile()

	return true
}

func GetHTTPClient() *http.Client {
	ctx := context.Background()

	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	return setup42.Client(ctx, token)
}
