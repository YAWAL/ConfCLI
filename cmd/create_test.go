package cmd

import (
	"reflect"
	"testing"
	"encoding/json"
	"github.com/YAWAL/GetMeConf/entities"
)

const (
	mongo      = "mongo.csv"
	tempConfig = "tempcnf.csv"
	tsConfig   = "tscnf.csv"
)

var mongoWantStruct = entities.Mongodb{Domain: "mongotest", Mongodb: true, Host: "mongo_test", Port: "mongoTest"}
var tempWantStruct = entities.Tempconfig{RestApiRoot: "tempconfigtest", Host: "tempconfig_test", Port: "tempConfigTest",
	Remoting: "tempconfigtest", LegasyExplorer: false}
var tsWantStruct = entities.Tsconfig{Module: "tscnftest", Target: "tscnf_test", SourceMap: false, Excluding: 11}

func Test_createByteConfig(t *testing.T) {
	mongoWant, _ := json.Marshal(mongoWantStruct)
	tempWant, _ := json.Marshal(tempWantStruct)
	tsWant, _ := json.Marshal(tsWantStruct)

	tests := []struct {
		name     string
		fileName string
		want     []byte
	}{
		{"mongo", mongo, mongoWant},
		{"temp", tempConfig, tempWant},
		{"ts", tsConfig, tsWant},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createByteConfig(tt.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test %v: createByteConfig() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
