package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "client functionality",
	Long:  `client functionality for GetMeConf service`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error during execution client: %v", err)
	}
}
