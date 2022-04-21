package jsonparser_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	jsonparser "github.com/Dysproz/ports-db-microservices/internal/core/services/jsonparse"
)

func TestStreamReading(t *testing.T) {
	stream := jsonparser.NewStream()
	go func() {
		for data := range stream.Watch() {
			require.NoError(t, data.Error)
			assert.Equal(t, testPorts[data.Key], data.Port)
		}
	}()
	pwd, _ := os.Getwd()
	stream.Load(pwd + "/testfiles/test_ports.json")
}

var testPorts = map[string]domain.Port{
	"AEAJM": {
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
	"AEAUH": {
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
	"AEDXB": {
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
