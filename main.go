package main

import (
	"fmt"
	"log"
	"main/commands"
	"main/ft_auth"
	"os"
	"text/tabwriter"

	"github.com/joho/godotenv"
)

func printHelp(cmds []commands.FtCommand) {
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)

	for _, cmd := range cmds {
		w_alias := ""
		for _, alias := range cmd.GetAlias() {
			if len(w_alias) > 0 {
				w_alias += ","
			}
			w_alias += alias
		}
		if len(w_alias) > 0 {
			w_alias = "Alias: " + w_alias
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t\n", cmd.GetCommand(), w_alias, cmd.GetDescription())

	}

	w.Flush()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if isLogged := ft_auth.Authenticate(os.Getenv("OAUTH_ID"), os.Getenv("OAUTH_SECRET")); !isLogged {
		log.Fatal("Can't retrieve the OAuth Token... Please, try again.")
	}

	cmds := LoadCommands()

	if len(os.Args) < 2 {
		printHelp(cmds)
		os.Exit(1)
	}

	for _, command := range cmds {
		command.Init()
	}

	for _, command := range cmds {
		if command.GetCommand() == os.Args[1] {
			command.GetFlags().Parse(os.Args[2:])
			command.Handler(os.Args[2:])
			return
		}
		for _, alias := range command.GetAlias() {
			if alias == os.Args[1] {
				command.GetFlags().Parse(os.Args[2:])
				command.Handler(os.Args[2:])
				return
			}
		}
	}

	printHelp(cmds)
	os.Exit(1)
}

func LoadCommands() []commands.FtCommand {
	cmds := make([]commands.FtCommand, 0)

	cmds = append(cmds, &commands.FtCommandMe{})
	cmds = append(cmds, &commands.FtCommandUser{})

	return cmds
}
