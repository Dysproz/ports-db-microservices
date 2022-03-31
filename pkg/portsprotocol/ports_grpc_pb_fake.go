package portsprotocol

import (
	context "context"
	"errors"

	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// FakePort is used for returning non-empty response from fakePortServiceClient.GetPort(_)
var FakePort = Port{
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

type fakePortServiceClient struct {
}

// NewFakePortServiceClient returns fakePortServiceClient interface implementation for testing purposes
func NewFakePortServiceClient() PortServiceClient {
	return &fakePortServiceClient{}
}

// CreateOrUpdatePort is used to return always empty response simulating success
func (c *fakePortServiceClient) CreateOrUpdatePort(ctx context.Context, in *CreateOrUpdatePortRequest, opts ...grpc.CallOption) (*CreateOrUpdatePortResponse, error) {
	return &CreateOrUpdatePortResponse{}, nil
}

// GetPort is used to return FakePort as fake response for testing purposes
func (c *fakePortServiceClient) GetPort(ctx context.Context, in *GetPortRequest, opts ...grpc.CallOption) (*GetPortResponse, error) {
	log.Info("GetPort ", in.Key, " port with fake grpc pb client")
	if in.Key == "fakePort" {
		log.Info("Returning fakePort with fake values.")
		return &GetPortResponse{
			Port: &FakePort,
		}, nil
	} else {
		log.Info("Returning empty fakePort.")
		return &GetPortResponse{}, errors.New("Port does not exist")
	}
}
