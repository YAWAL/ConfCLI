package cmd

import (
	"encoding/csv"
	"os"
	"log"
	"strconv"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/YAWAL/GetMeConfAPI/api"
	"github.com/YAWAL/GetMeConf/entitie"
)

const trueRecord = "true"
const falseRecord = "false"

func readConfig(fileName string) ([][]string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("Error during opening file has occurred: %v", err)
		return nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error during reading file has occurred: %v", err)
		return nil
	}
	return records
}

func createByteConfig(fileName string) []byte {
	records := readConfig(fileName)
	switch fileName {
	case "mongo.csv":
		var mongocnf entitie.Mongodb
		mongocnf.Domain = records[0][0]
		if records[0][1] != trueRecord && records[0][1] != falseRecord {
			log.Fatalf("Field Mongodb should be true or false, but is: %v", records[0][1])
			return nil
		}
		if records[0][1] == trueRecord {
			mongocnf.Mongodb = true
		}
		if records[0][1] == falseRecord {
			mongocnf.Mongodb = false
		}
		mongocnf.Host = records[0][2]
		mongocnf.Port = records[0][3]
		bytesMongo, err := json.Marshal(mongocnf)
		if err != nil {
			log.Printf("Error during converting Mongodb structure to []byte has occurred: %v", err)
		}
		return bytesMongo
	case "tempcnf.csv":
		var tempcnf entitie.Tempconfig
		tempcnf.RestApiRoot = records[0][0]
		tempcnf.Host = records[0][1]
		tempcnf.Port = records[0][2]
		tempcnf.Remoting = records[0][3]
		if records[0][4] != trueRecord && records[0][4] != falseRecord {
			log.Fatalf("Field legasyExplorer should be true or false, but is: %v", records[0][4])
			return nil
		}
		if records[0][4] == trueRecord {
			tempcnf.LegasyExplorer = true
		}
		if records[0][4] == falseRecord {
			tempcnf.LegasyExplorer = false
		}
		bytesTempcnf, err := json.Marshal(tempcnf);
		if err != nil {
			log.Printf("Error during converting Tempconfig structure to []byte has occurred: %v", err)
		}
		return bytesTempcnf
	case "tscnf.csv":
		var tscnf entitie.Tsconfig
		tscnf.Module = records[0][0]
		tscnf.Target = records[0][1]
		if records[0][2] != trueRecord && records[0][2] != falseRecord {
			log.Printf("Field sourceMap should be true or false, but is: %v", records[0][2])
			return nil
		}
		if records[0][2] == trueRecord {
			tscnf.SourceMap = true
		}
		if records[0][2] == falseRecord {
			tscnf.SourceMap = false
		}
		excluding, err := strconv.Atoi(records[0][3])
		if err != nil {
			log.Printf("field Excluding should be integer, but is: %T", records[0][3])
		}
		tscnf.Excluding = excluding
		bytesTscnf, err := json.Marshal(tscnf)
		if err != nil {
			log.Printf("Error during converting Tsconfig structure to []byte has occurred: %v", err)
		}
		return bytesTscnf
	default:
		log.Printf("Cant find file: %v", fileName)
	}
	return nil
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create config command",
	Long:  `create is a command for creating config and persist it to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		log.Printf("State: %v", conn.GetState())
		if err != nil {
			log.Fatalf("Dial error has occurred: %v", err)
		}
		defer conn.Close()
		client := api.NewConfigServiceClient(conn)
		config := createByteConfig(fileName)
		resp, err := client.CreateConfig(context.Background(), &api.Config{Config: config, ConfigType: configType})
		if err != nil {
			log.Printf("State: %v", conn.GetState())
			log.Printf("Error during client.CreateConfig has occurred: %v", err)
		}
		if resp.Status != "OK" {
			log.Printf("State: %v", conn.GetState())
			log.Printf("Error during creating config has occurred: %v response status: %v", err, resp.Status)
		}
		log.Printf("State: %v", conn.GetState())
	},
}
