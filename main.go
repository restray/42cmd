package main

import (
	"flag"
	"fmt"
	"log"
	"main/ft_auth"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if isLogged := ft_auth.Authenticate(os.Getenv("OAUTH_ID"), os.Getenv("OAUTH_SECRET")); !isLogged {
		log.Fatal("Can't retrieve the OAuth Token... Please, try again.")
	}

	// fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// fooEnable := fooCmd.Bool("enable", false, "enable")
	// fooName := fooCmd.String("name", "", "name")

	// barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	// barLevel := barCmd.Int("level", 0, "level")

	meCmd := flag.NewFlagSet("me", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	// case "foo":
	// 	fooCmd.Parse(os.Args[2:])
	// 	fmt.Println("subcommand 'foo'")
	// 	fmt.Println("  enable:", *fooEnable)
	// 	fmt.Println("  name:", *fooName)
	// 	fmt.Println("  tail:", fooCmd.Args())
	// case "bar":
	// 	barCmd.Parse(os.Args[2:])
	// 	fmt.Println("subcommand 'bar'")
	// 	fmt.Println("  level:", *barLevel)
	// 	fmt.Println("  tail:", barCmd.Args())
	case "me":
		meCmd.Parse(os.Args[2:])
		retrieveMe()
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
