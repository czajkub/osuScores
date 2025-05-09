package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	
	"time"
	"github.com/spf13/cobra"

	"github.com/spf13/cobra"

	"github.com/joho/godotenv"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get most recent 1000 scores",
	Long:  "Returns JSON-style string containing the 1000 most recent osu! scores with the gamemode specified by the -g flag",
	Run: func(cmd *cobra.Command, args []string) {
		getscores()
	},
}

func getscores() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	token := os.Getenv("ACCESS_TOKEN")
	if token == "" {
		fmt.Println("Access token not added to .env file!")
		Startoauth()
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/v2/scores?ruleset=osu", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)

}
