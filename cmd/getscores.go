package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"slices"

	"github.com/joho/godotenv"
	"time"
)

func Getscores() []byte {
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

	if resp.StatusCode != 200 {
		fmt.Printf("Got response status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}
	return body
}

func Sortscores(scores []Score, bywhat string) []Score {
	switch bywhat {
	case "pp":
		slices.SortFunc(scores, ppcomp)
		return scores
	case "score":
		slices.SortFunc(scores, scorecomp)
		return scores
	case "acc":
		slices.SortFunc(scores, acccomp)
		return scores
	}
	return scores
}

func ppcomp(lhs, rhs Score) int {
	if lhs.Pp < rhs.Pp {
		return -1
	} else if lhs.Pp > rhs.Pp {
		return 1
	} else {
		return 0
	}
}
func scorecomp(lhs, rhs Score) int {
	if lhs.Total_score < rhs.Total_score {
		return -1
	} else if lhs.Total_score > rhs.Total_score {
		return 1
	}
	return ppcomp(lhs, rhs)
}
func acccomp(lhs, rhs Score) int {
	if lhs.Accuracy < rhs.Accuracy {
		return -1
	} else if lhs.Accuracy > rhs.Accuracy {
		return 1
	}
	return ppcomp(lhs, rhs)
}
