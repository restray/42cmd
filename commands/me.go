package commands

import (
	"main/intrapi"

	"encoding/json"
	"flag"
	"fmt"
)

type FtCommandMe struct{}

func (me *FtCommandMe) GetCommand() string {
	return "me"
}

func (me *FtCommandMe) GetAlias() []string {
	return []string{}
}

func (me *FtCommandMe) GetDescription() string {
	return "Give your 42 profile"
}

func (me *FtCommandMe) Init() {
	flag.NewFlagSet("me", flag.ExitOnError)
}

func (me *FtCommandMe) Handler(args []string) {
	result := intrapi.MakeApiRequest("/me")

	var user intrapi.User42
	json.Unmarshal(result, &user)

	fmt.Println()
	fmt.Printf(`	██╗  ██╗██████╗ 
	██║  ██║╚════██╗	%s
	███████║ █████╔╝	%s
	╚════██║██╔═══╝ 	Level %.0f
	     ██║███████╗	lvl 0 [%s] lvl 21
	     ╚═╝╚══════╝`, user.UsualFullName, user.Login, user.CursusUsers[1].Level, loadingBar(0, 21, int(user.CursusUsers[1].Level)))
	fmt.Println()
	fmt.Println()
}
