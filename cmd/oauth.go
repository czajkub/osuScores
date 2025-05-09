package cmd

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"time"
)

func loadenv() [2]string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return [2]string{"", ""}
	}
	client_id := os.Getenv("CLIENT_ID")
	if client_id == "" {
		fmt.Println("Client ID environment variable is not set")
	}
	client_secret := os.Getenv("CLIENT_SECRET")
	if client_secret == "" {
		fmt.Println("client_secret environment variable is not set")
	}
	return [2]string{client_id, client_secret}
}

func generateStateString() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", b)
}

func getoauthconfig(client_id string, client_secret string) *oauth2.Config {
	//var authurl = "https://osu.ppy.sh/oauth/authorize?client_id=" + client_id + "&redirect_uri=http%3A%2F%2Flocalhost%3A8080&response_type=code&scope=public+identify"
	//var tokenurl = "https://osu.ppy.sh/oauth/token?client_id=" + client_id + "&client_secret=" + client_secret + "&code=receivedcode&grant_type=authorization_code&redirect_uri=http%3A%2F%2Flocalhost%3A8080"

	var osuEndpoint = oauth2.Endpoint{
		AuthURL:  "https://osu.ppy.sh/oauth/authorize",
		TokenURL: "https://osu.ppy.sh/oauth/token",
	}

	var oauth2Config = oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"public", "identify"},
		Endpoint:     osuEndpoint,
	}

	return &oauth2Config
}

func Startoauth() {
	envvars := loadenv()
	client_id := envvars[0]
	client_secret := envvars[1]
	if client_id == "" || client_secret == "" {
		fmt.Println("Client ID and Client Secret are required")
		os.Exit(1)
	}

	//fmt.Println(client_id, client_secret)
	state := generateStateString()

	oauthConfig := getoauthconfig(client_id, client_secret)
	authURL := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: \n%v\n", authURL)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		// Validate the state parameter.
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State mismatch", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")
		token, err := oauthConfig.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error exchanging code for token: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Println("Add this token to your .env file")
		fmt.Println(token.AccessToken)

		fmt.Fprintln(w, "Authentication successful! You can close this window.")
		go func() {
			time.Sleep(2 * time.Second)
			os.Exit(0)
		}()

	})

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
