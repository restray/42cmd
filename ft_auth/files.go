package ft_auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/oauth2"
)

func retrieveTokenFromFile() *oauth2.Token {
	var retrieved_token *oauth2.Token

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	if file, err := ioutil.ReadFile(path.Join(exPath, "token.json")); err == nil {
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
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	ioutil.WriteFile(path.Join(exPath, "token.json"), resp, 0600)
}
