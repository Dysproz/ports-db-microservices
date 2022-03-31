package resthandler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Dysproz/ports-db-microservices/pkg/grpc"
	"github.com/Dysproz/ports-db-microservices/pkg/jsonparser"
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

// PortRequest struct is used to decode into getPort json request
type PortRequest struct {
	Key string `json:"key"`
}

// RESTClient handles incoming HTTP requests
type RESTClient struct {
	Client pb.PortServiceClient
	Stream jsonparser.Stream
}

// HandleRequests method handles incoming HTTP requests and routes logic.
func HandleRequests(client pb.PortServiceClient, stream jsonparser.Stream) {
	restClient := RESTClient{Client: client, Stream: stream}
	http.HandleFunc("/getPort", restClient.HandleGetPort)
	http.HandleFunc("/loadPorts", restClient.HandleLoadPorts)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// HandleGetPort handles HTTP requests for getPort route
func (c *RESTClient) HandleGetPort(w http.ResponseWriter, r *http.Request) {
	var jsonRequest PortRequest
	if err := json.NewDecoder(r.Body).Decode(&jsonRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Debug("Got getPort request for key: ", jsonRequest.Key)
	retrievedPort, err := grpc.GetPort(c.Client, jsonRequest.Key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonRetrievedPort, err := protojson.Marshal(&retrievedPort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Debug("Successfully found port for key ", jsonRequest.Key, " with port name: ", retrievedPort.Name)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(jsonRetrievedPort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// HandleLoadPorts handles HTTP requests for loadPorts route
func (c *RESTClient) HandleLoadPorts(w http.ResponseWriter, r *http.Request) {
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
	go c.Stream.Start(f.Name())
	w.WriteHeader(http.StatusOK)
}
