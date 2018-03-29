package cmd

import (
	"log"
	"google.golang.org/grpc"
	"github.com/YAWAL/GetMeConfAPI/api"
	"golang.org/x/net/context"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete config",
	Long:  `Delete config from database`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		log.Printf("State: %v", conn.GetState())
		if err != nil {
			log.Printf("State: %v", conn.GetState())
			log.Fatalf("Dial error has occurred: %v", err)
		}
		defer conn.Close()
		client := api.NewConfigServiceClient(conn)
		resp, err := client.DeleteConfig(context.Background(), &api.DeleteConfigRequest{ConfigType: configType, ConfigName: configName})
		if err != nil {
			log.Printf("State: %v", conn.GetState())
			log.Printf("Error during client.DeleteConfig has occurred: %v", err)
		}
		if resp.Status == "" {
			log.Printf("State: %v", conn.GetState())
			log.Printf("Response status is empty: %v", resp.Status)
		}
		log.Printf("Config name: %v, cofig type: %v\n", configName, configType)
		log.Printf("Response status: %v", resp.Status)
		log.Printf("State: %v", conn.GetState())
	},
}
