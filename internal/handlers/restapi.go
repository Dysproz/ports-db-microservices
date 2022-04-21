package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
	"github.com/Dysproz/ports-db-microservices/internal/core/ports"
)

type restClient struct {
	Client       ports.GRPCClientService
	parseService ports.JSONParseService
}

type portRequest struct {
	Key string `json:"key"`
}

// NewRESTClient return new REST Client
func NewRESTClient(portServiceClient ports.GRPCClientService, parseService ports.JSONParseService) *restClient {
	return &restClient{
		Client:       portServiceClient,
		parseService: parseService,
	}
}

// HandleRequests method handles incoming HTTP requests and routes logic.
func (c *restClient) HandleRequests() {
	http.HandleFunc("/getPort", c.HandleGetPort)
	http.HandleFunc("/loadPorts", c.HandleLoadPorts)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// HandleGetPort Handles HTTP request for getting port
func (c *restClient) HandleGetPort(w http.ResponseWriter, r *http.Request) {
	var jsonRequest portRequest
	if err := json.NewDecoder(r.Body).Decode(&jsonRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Debug("Got getPort request for key: ", jsonRequest.Key)
	retrievedPort, err := c.Client.GetPort(context.Background(), &domain.GetPortRequest{
		Key: jsonRequest.Key,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonRetrievedPort, err := protojson.Marshal(retrievedPort.Port)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Debug("Successfully found port for key ", jsonRequest.Key, " with port name: ", retrievedPort.Port.Name)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(jsonRetrievedPort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

// HandleLoadPorts handles HTTP requests to load ports from file
func (c *restClient) HandleLoadPorts(w http.ResponseWriter, r *http.Request) {
	log.Debug("Got loadPorts request")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Debug("Failed on loading file")
		log.Debug(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer file.Close()
	f, err := os.OpenFile("/root/ports.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Debug("Failed on opening file")
		log.Debug(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	io.Copy(f, file)
	go c.parseService.Load(f.Name())
	w.WriteHeader(http.StatusOK)
}
