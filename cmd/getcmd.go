package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get most recent 1000 scores",
	Long:  "Prints data from 1000 most recent osu! scores with the gamemode specified by the -g flag",
	Run: func(cmd *cobra.Command, args []string) {
		rawscores := Getscores()
		scores := UnmarshalJSON(rawscores)
		printscores(scores)
	},
}

func printscores(scores []Score) {
	for idx, score := range scores {
		fmt.Printf("Score number %d\n", idx+1)
		fmt.Printf("pp: %f; accuracy: %f; rank: %s; userid: %d\n", score.Pp, score.Accuracy, score.Rank, score.User_id)
	}
}
