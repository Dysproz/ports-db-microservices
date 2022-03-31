package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/Dysproz/ports-db-microservices/pkg/mongodb"
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
	"github.com/Dysproz/ports-db-microservices/pkg/portsprotocolserver"
)

func main() {
	defer os.Exit(0)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	port, mongodbAddress, err := getParameters()
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	mongodbClient, err := mongodb.CreateMongoDBClient(mongodbAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer mongodbClient.Close()
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	defer grpcServer.GracefulStop()
	pb.RegisterPortServiceServer(grpcServer, newServer(mongodbClient))
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	select {
	case <-sigCh:
		log.Info("Interrupt signal detected. Gracefully shutting down...")
		runtime.Goexit()
	}
}

func newServer(mongo mongodb.MongoClient) *portsprotocolserver.PortsProtocolServer {
	s := &portsprotocolserver.PortsProtocolServer{
		MongoDB: mongo,
	}
	return s
}

func getParameters() (int, string, error) {
	port := flag.Int("port", 5000, "Server port")
	mongodbAddres := flag.String("mongodb-address", "mongodb_1:27017", "Address of mongoDB database")
	logLevel := flag.String("log-level", "info", "Log level.")
	flag.Parse()
	logLevelParsed, err := log.ParseLevel(*logLevel)
	if err != nil {
		return 0, "", err
	}
	log.SetLevel(logLevelParsed)
	log.Debug("Running with CLI parameters. port: ", *port,
		" mongodb-address: ", *mongodbAddres,
		" log-level: ", *logLevel)
	return *port, *mongodbAddres, nil
}
