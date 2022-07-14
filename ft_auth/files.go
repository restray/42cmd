package ft_auth

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"

	"golang.org/x/oauth2"
)

func retrieveTokenFromFile() *oauth2.Token {
	var retrieved_token *oauth2.Token

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFolderPath := path.Join(dirname, ".42cmd")

	if file, err := ioutil.ReadFile(path.Join(configFolderPath, "token.json")); err == nil {
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

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFolderPath := path.Join(dirname, ".42cmd")

	ioutil.WriteFile(path.Join(configFolderPath, "token.json"), resp, fs.FileMode(int(0600)))
}
