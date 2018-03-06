package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	configName string
	configType string
	outPath    string
	fileName   string
	address    string
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "client functionality",
	Long:  `client functionality for GetMeConf service`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configName, "config-name", "", "config name")
	rootCmd.PersistentFlags().StringVar(&configType, "config-type", "", "config type")
	rootCmd.PersistentFlags().StringVar(&outPath, "outpath", "", "output path for config file")
	rootCmd.PersistentFlags().StringVar(&fileName, "file-name", "", "config file's name")
	rootCmd.PersistentFlags().StringVar(&address, "address", "localhost:3000", "server address")
	rootCmd.AddCommand(createCmd, deleteCmd, readCmd, updateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error during execution client: %v", err)
	}
}
