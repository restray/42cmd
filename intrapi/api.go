package intrapi

import (
	"io/ioutil"
	"log"
	"main/ft_auth"
)

func makeAPIReq(req string) []byte {
	client := ft_auth.GetHTTPClient()

	response, err := client.Get("https://api.intra.42.fr/v2" + req)
	if err != nil {
		log.Fatalf("failed getting user info: %s\n", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed read response: %s\n", err.Error())
	}
	return contents
}
