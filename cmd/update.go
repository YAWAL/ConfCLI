package cmd

import (
	"log"
	"google.golang.org/grpc"
	"github.com/YAWAL/GetMeConfAPI/api"
	"golang.org/x/net/context"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update config",
	Long:  `Update existing config in database`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		log.Printf("State: %v", conn.GetState())
		if err != nil {
			log.Fatalf("Dial error has occurred: %v", err)
		}
		defer conn.Close()
		client := api.NewConfigServiceClient(conn)
		config := createByteConfig(fileName)
		resp, err := client.UpdateConfig(context.Background(), &api.Config{Config: config, ConfigType: configType})
		if err != nil {
			log.Printf("Error during client.CreateConfig has occurred: %v", err)
		}
		if resp.Status != "OK" {
			log.Printf("Error during creating config has occurred: %v response status: %v", err, resp.Status)
		}
		log.Printf("Response: %v", resp.Status)
	},
}
