package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var setup42 *oauth2.Config

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	setup42 = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_ID"),
		ClientSecret: os.Getenv("OAUTH_SECRET"),
		Scopes:       []string{"public", "projects", "profile", "tig"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.intra.42.fr/oauth/authorize",
			TokenURL: "https://api.intra.42.fr/oauth/token",
		},
		RedirectURL: "http://localhost:25634/callback",
	}

	mux.HandleFunc("/", oauthGoogleLogin)
	mux.HandleFunc("/callback", oauthGoogleCallback)

	return mux
}

func getHTTPClient() *http.Client {
	ctx := context.Background()

	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	return setup42.Client(ctx, Token)
}

func oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	u := setup42.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	err := retrieveToken(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Connecté\n")
	Server.Close()
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func retrieveToken(code string) error {
	var err error
	Token, err = setup42.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	return nil
}
