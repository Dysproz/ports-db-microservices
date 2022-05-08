package transport

import (
	context "context"
	"errors"

	domain "github.com/Dysproz/ports-db-microservices/internal/core/domain"
	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// FakePorts is a set of fake ports used for tests
var FakePorts = map[string]*domain.Port{
	"fakePort": {
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
	},
	"fakePort2": {
		Name:        "fakeName2",
		City:        "fakeCity2",
		Country:     "fakeCountry2",
		Alias:       []string{},
		Regions:     []string{"fakeregion2"},
		Coordinates: []float32{33, 44},
		Province:    "fakeProvince2",
		Timezone:    "fakeTimezone2",
		Unlocs:      []string{},
		Code:        "fakeCode2",
	},
	"fakePort3": {
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
	},
}

// FakePort is used for returning non-empty response from fakePortServiceClient.GetPort(_)
var FakePort = domain.Port{
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
func (c *fakePortServiceClient) CreateOrUpdatePort(ctx context.Context, in *domain.CreateOrUpdatePortRequest, opts ...grpc.CallOption) (*domain.CreateOrUpdatePortResponse, error) {
	return &domain.CreateOrUpdatePortResponse{}, nil
}

// GetPort is used to return FakePort as fake response for testing purposes
func (c *fakePortServiceClient) GetPort(ctx context.Context, in *domain.GetPortRequest, opts ...grpc.CallOption) (*domain.GetPortResponse, error) {
	log.Info("GetPort ", in.Key, " port with fake grpc pb client")
	if _, ok := FakePorts[in.Key]; ok {
		log.Info("Returning fakePort with fake values.")
		return &domain.GetPortResponse{
			Port: FakePorts[in.Key],
		}, nil
	}
	log.Info("Returning empty fakePort.")
	return &domain.GetPortResponse{}, errors.New("Port does not exist")
}
