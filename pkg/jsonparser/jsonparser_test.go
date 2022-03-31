package jsonparser_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Dysproz/ports-db-microservices/pkg/jsonparser"
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

func TestStreamReading(t *testing.T) {
	stream := jsonparser.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			require.NoError(t, data.Error)
			assert.Equal(t, testPorts[data.Key], data.Port)
		}
	}()
	pwd, _ := os.Getwd()
	stream.Start(pwd + "/testfiles/test+ports.json")
}

var testPorts = map[string]pb.Port{
	"AEAJM": pb.Port{
		Name:    "Ajman",
		City:    "Ajman",
		Country: "United Arab Emirates",
		Alias:   []string{},
		Regions: []string{},
		Coordinates: []float32{
			55.5136433,
			25.4052165,
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEAJM",
		},
		Code: "52000",
	},
	"AEAUH": pb.Port{
		Name:    "Abu Dhabi",
		City:    "Abu Dhabi",
		Country: "United Arab Emirates",
		Alias:   []string{},
		Regions: []string{},
		Coordinates: []float32{
			54.37,
			24.47,
		},
		Province: "Abu Z¸aby [Abu Dhabi]",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEAUH",
		},
		Code: "52001",
	},
	"AEDXB": pb.Port{
		Name:    "Dubai",
		City:    "Dubai",
		Country: "United Arab Emirates",
		Alias:   []string{},
		Regions: []string{},
		Coordinates: []float32{
			55.27,
			25.25,
		},
		Province: "Dubayy [Dubai]",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEDXB",
		},
		Code: "52005",
	},
}
