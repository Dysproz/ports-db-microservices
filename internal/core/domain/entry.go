package domain

import (
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

type Entry struct {
	Error error
	Key   string
	Port  pb.Port
}
