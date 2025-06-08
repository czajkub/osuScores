package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type results struct {
	Scores        []Score `json:"scores"`
	Cursor        Cursor  `json:"cursor"`
	Cursor_string string  `json:"cursor_string"`
}

type Cursor struct {
	Id int `json:"id"`
}
type Score struct {
	Accuracy            float32       `json:"accuracy"`
	Beatmap_id          int           `json:"beatmap_id"`
	Best_id             int           `json:"best_id"`
	Build_id            int           `json:"build_id"`
	Classic_total_score int           `json:"classic_total_score"`
	Ended_at            string        `json:"ended_at"`
	Has_replay          bool          `json:"has_replay"`
	Id                  int           `json:"id"`
	Is_perfect_combo    bool          `json:"is_perfect_combo"`
	Legacy_perfect      bool          `json:"legacy_perfect"`
	Legacy_score_id     int           `json:"legacy_score_id"`
	Legacy_total_score  int           `json:"legacy_total_score"`
	Max_combo           int           `json:"max_combo"`
	maximum_statistics  maxstatistics `json:"maximum_statistics"`
	Mods                []modsused    `json:"mods"`
	Passed              bool          `json:"passed"`
	Playlist_item_id    int           `json:"playlist_item_id"`
	Pp                  float32       `json:"pp"`
	Preserve            bool          `json:"preserve"`
	Processed           bool          `json:"processed"`
	Rank                string        `json:"rank"`
	Ranked              bool          `json:"ranked"`
	Room_id             int           `json:"room_id"`
	Ruleset_id          int           `json:"ruleset_id"`
	Started_at          string        `json:"started_at"`
	Statistics          maxstatistics `json:"statistics"`
	Total_score         int           `json:"total_score"`
	Scoretype           string        `json:"type"`
	User_id             int           `json:"user_id"`
}

type maxstatistics struct {
	Great                 int `json:"great"`
	Legacy_combo_increase int `json:"legacy_combo_increase"`
}
type modsused struct {
	Modacronym string `json:"acronym"`
}

func UnmarshalJSON(data []byte) []Score {
	var unmarshaledJSON results
	err := json.Unmarshal(data, &unmarshaledJSON)
	if err != nil {
		fmt.Println("Error unmarshaling JSON: ", err)
		os.Exit(1)
	}
	fmt.Printf("1st score pp: %f\n", unmarshaledJSON.Scores[0].Pp)
	return unmarshaledJSON.Scores
}
