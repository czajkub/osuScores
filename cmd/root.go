package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "osuscores",
	Short: "Get up to 1000 recent osu! scores",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,

	//Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.main.yaml)")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	viper.SetDefault("author", "czajkub <czajkub@gmail.com>")
	viper.SetDefault("license", "gpl-3.0")

	getCmd.Flags().StringVarP(&Comparable, "compare", "c", "pp", "value to compare scores by")
	rootCmd.AddCommand(getCmd)

}

func initConfig() {

	// If a config file is found, read it in.
	if cfgFile := viper.GetString("config"); cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			return
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".main")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
