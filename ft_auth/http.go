package ft_auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func (s *server) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}

	log.Printf("Stopping http server ...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}

func openNavigator(s *server) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://"+s.Addr).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://"+s.Addr).Start()
	case "darwin":
		err = exec.Command("open", "http://"+s.Addr).Start()
	default:
		err = fmt.Errorf("Your OS is not supported. Visit http://" + s.Addr)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func serveToken() {
	server := &server{
		Server: http.Server{
			Addr:         fmt.Sprintf("localhost:25634"),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", oauth42Login)
	mux.HandleFunc("/callback", server.oauth42Callback)
	server.Server.Handler = mux

	done := make(chan bool)
	go func() {
		defer server.Server.Close()
		if err := server.Server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("%v", err)
		}
		done <- true
	}()

	openNavigator(server)

	server.WaitShutdown()
}

func oauth42Login(w http.ResponseWriter, r *http.Request) {
	oauthState := generateStateOauthCookie(w)

	u := setup42.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (s *server) oauth42Callback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth state token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	err := retrieveToken(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "You're logged in, go back to your terminal.")

	go func() {
		s.shutdownReq <- true
	}()

	return
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

	token, err = setup42.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	return nil
}
