package cmd

import "github.com/spf13/cobra"

var modCmd = &cobra.Command{
	Use:   "mods",
	Short: "Most commonly used mods",
	Long:  "Prints most commonly used mods in the past 1000 osu! plays",
	Run: func(cmd *cobra.Command, args []string) {
		rawscores := Getscores()
		scores := UnmarshalJSON(rawscores)
		Usedmods(scores)
	},
}
