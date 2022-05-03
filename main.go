package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var Server *http.Server
var Token *oauth2.Token

func ServeToken() {
	// We create a simple server using http.Server and run.
	Server = &http.Server{
		Addr:    fmt.Sprintf("localhost:25634"),
		Handler: NewHandler(),
	}

	defer Server.Close()

	log.Printf("Starting HTTP Server. Listening at %q", Server.Addr)

	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://"+Server.Addr).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://"+Server.Addr).Start()
	case "darwin":
		err = exec.Command("open", "http://"+Server.Addr).Start()
	default:
		err = fmt.Errorf("Your OS is not supported. Visit http://" + Server.Addr)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err := Server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	for Token == nil || !Token.Valid() {
		ServeToken()
	}

	/** @TODO Need to save the json in config file */
	// resp, err := json.Marshal(Token)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	meCmd := flag.NewFlagSet("me", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	case "me":
		meCmd.Parse(os.Args[2:])
		retrieveMe()
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
