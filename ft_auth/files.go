package ft_auth

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
)

func retrieveTokenFromFile() *oauth2.Token {
	var retrieved_token *oauth2.Token

	if file, err := ioutil.ReadFile("token.json"); err == nil {
		json.Unmarshal([]byte(file), &retrieved_token)
	}

	if !retrieved_token.Valid() {
		return nil
	}

	return retrieved_token
}

func saveTokenOnFile() {
	resp, err := json.Marshal(token)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile("token.json", resp, 0600)
}
