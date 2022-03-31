package resthandler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/Dysproz/ports-db-microservices/pkg/jsonparser"
	"github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
	"github.com/Dysproz/ports-db-microservices/pkg/resthandler"
	"github.com/stretchr/testify/assert"
)

func TestHandleGetPort(t *testing.T) {
	client := resthandler.RESTClient{
		Client: portsprotocol.NewFakePortServiceClient(),
		Stream: jsonparser.NewJSONStream(),
	}
	var jsonPayload = []byte(`{"key":"fakePort"}`)
	req := httptest.NewRequest("POST", "http://example.com/getPort", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(client.HandleGetPort)
	handler.ServeHTTP(w, req)
	assert.NotEqual(t, w.Code, http.StatusInternalServerError)
}

var testPort = pb.Port{
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

func TestHandleLoadPorts(t *testing.T) {
	client := resthandler.RESTClient{
		Client: portsprotocol.NewFakePortServiceClient(),
		Stream: jsonparser.NewJSONStream(),
	}
	req := httptest.NewRequest("POST", "http://example.com/getPort", bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(client.HandleLoadPorts)
	handler.ServeHTTP(w, req)
	assert.NotEqual(t, w.Code, http.StatusInternalServerError)
}
