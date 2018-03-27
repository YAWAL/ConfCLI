package cmd

import (
	"log"
	"io"
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/YAWAL/GetMeConfAPI/api"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
	"github.com/YAWAL/GetMeConf/entitie"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read config",
	Long:  `read configs from database`,
	Run: func(cmd *cobra.Command, args []string) {
		if configName == "" && configType == "" {
			log.Fatal("Can't process => config name and config type are empty")
		}
		log.Printf("Start checking input data:\n Config name: %v\n Config type : %v\n Output path: %v\nProcessing ...", configName, configType, outPath)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		log.Printf("State: %v", conn.GetState())
		if err != nil {
			log.Printf("State: %v", conn.GetState())
			log.Fatalf("Dial error has occurred: %v", err)
		}
		defer conn.Close()
		client := api.NewConfigServiceClient(conn)
		if configName != "" && configType != "" {
			log.Printf("State: %v", conn.GetState())
			log.Println("Processing retrieving config...")
			err := retrieveConfig(configName, outPath, client)
			if err != nil {
				log.Printf("State: %v", conn.GetState())
				log.Fatalf("retrieveConfig err: %v", err)
			}
		}
		if configName == "" && configType != "" {
			err := retrieveConfigs(client)
			if err != nil {
				log.Printf("State: %v", conn.GetState())
				log.Fatalf("retrieveConfigs err : %v", err)
			}
		}
		log.Println("End retrieving config.")
		log.Printf("State: %v", conn.GetState())
	},
}

func retrieveConfig(fileName, outputPath string, client api.ConfigServiceClient) error {
	conf, err := client.GetConfigByName(context.Background(), &api.GetConfigByNameRequest{ConfigName: configName, ConfigType: configType})
	if err != nil {
		log.Fatalf("Error during retrieving config has occurred: %v", err)
		return err
	}
	if err := writeFile(conf.Config, fileName, outputPath); err != nil {
		log.Fatalf("Error during writing file in retrieving config: %v", err)
		return err
	}
	return nil
}

func retrieveConfigs(client api.ConfigServiceClient) error {
	stream, err := client.GetConfigsByType(context.Background(), &api.GetConfigsByTypeRequest{ConfigType: configType})
	if err != nil {
		log.Fatalf("Error during retrieving stream configs has occurred:%v", err)
	}
	for {
		config, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error during streaming has occurred: %v", err)
			return err
		}
		switch configType {
		case "mongodb":
			var mongodb entitie.Mongodb
			err := json.Unmarshal(config.Config, &mongodb)
			if err != nil {
				log.Fatalf("Unmarshal mongodb err: %v", err)
			}
			flName := mongodb.Domain
			writeFile(config.Config, flName, outPath)
		case "tempconfig":
			var tempconfig entitie.Tempconfig
			err := json.Unmarshal(config.Config, &tempconfig)
			if err != nil {
				log.Fatalf("Unmarshal tempconfig err: %v", err)
			}
			flName := tempconfig.RestApiRoot
			writeFile(config.Config, flName, outPath)
		case "tsconfig":
			var tsconfig entitie.Tsconfig
			err := json.Unmarshal(config.Config, &tsconfig)
			if err != nil {
				log.Fatalf("Unmarshal tsconfig err: %v", err)
			}
			flName := tsconfig.Module
			writeFile(config.Config, flName, outPath)
		default:
			log.Fatalf("Config: %v does not exist", configType)
		}
	}
	return nil
}
