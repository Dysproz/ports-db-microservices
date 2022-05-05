package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	jsonparser "github.com/Dysproz/ports-db-microservices/internal/core/services/jsonparse"
	"github.com/Dysproz/ports-db-microservices/internal/core/services/portsprotocol"
	"github.com/Dysproz/ports-db-microservices/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type getPortTest struct {
	key          string
	expectedPort domain.Port
}

var testPort = domain.Port{
	Name:        "fakeName",
	City:        "fakeCity",
	Country:     "fakeCountry",
	Alias:       []string{"fakeAlias"},
	Regions:     []string{"fakeregion"},
	Coordinates: []float32{11.111, 22.222},
	Province:    "fakeProvince",
	Timezone:    "fakeTimezone",
	Unlocs:      []string{"fakeUnlock"},
	Code:        "fakeCode",
}

var testPort2 = domain.Port{
	Name:        "fakeName2",
	City:        "fakeCity2",
	Country:     "fakeCountry2",
	Alias:       nil,
	Regions:     []string{"fakeregion2"},
	Coordinates: []float32{33, 44},
	Province:    "fakeProvince2",
	Timezone:    "fakeTimezone2",
	Unlocs:      nil,
	Code:        "fakeCode2",
}

var testPort3 = domain.Port{
	Name:        "fakeName3",
	City:        "fakeCity2",
	Country:     "fakeCountry2",
	Alias:       []string{"fakeAliasX", "fakeAliasY"},
	Regions:     []string{"fakeregionX", "fakeRegionY"},
	Coordinates: []float32{12.765, 22.908},
	Province:    "fakeProvince3",
	Timezone:    "fakeTimezone3",
	Unlocs:      []string{"fakeUnlock", "fakeUnlock2"},
	Code:        "fakeCode3",
}

func TestHandleGetPort(t *testing.T) {
	tests := []getPortTest{
		{key: "fakePort", expectedPort: testPort},
		{key: "fakePort2", expectedPort: testPort2},
		{key: "fakePort3", expectedPort: testPort3},
	}
	client := handlers.NewRESTClient(
		portsprotocol.NewFakePortServiceClient(),
		jsonparser.NewStream(),
	)
	for _, tc := range tests {
		var jsonPayload = []byte(fmt.Sprintf(`{"key":"%v"}`, tc.key))
		req := httptest.NewRequest("POST", "http://example.com/getPort", bytes.NewBuffer(jsonPayload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(client.HandleGetPort)
		handler.ServeHTTP(w, req)
		assert.NotEqual(t, w.Code, http.StatusInternalServerError)
		var resultPort domain.Port
		require.NoError(t, json.NewDecoder(w.Result().Body).Decode(&resultPort))
		assert.Equal(t, resultPort, tc.expectedPort)
	}
}

func TestHandleLoadPorts(t *testing.T) {
	client := handlers.NewRESTClient(
		portsprotocol.NewFakePortServiceClient(),
		jsonparser.NewStream(),
	)
	req := httptest.NewRequest("POST", "http://example.com/getPort", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(client.HandleLoadPorts)
	handler.ServeHTTP(w, req)
	assert.NotEqual(t, w.Code, http.StatusInternalServerError)
}
