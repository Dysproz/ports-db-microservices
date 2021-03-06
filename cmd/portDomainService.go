package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/Dysproz/ports-db-microservices/internal/grpc"
	"github.com/Dysproz/ports-db-microservices/internal/repositories/portsrepo"
)

func main() {
	defer func() {
		log.Info("Port Domain Service fully stopped")
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-sigCh
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	port, mongodbAddress, err := getParameters()
	if err != nil {
		log.Error(err)
		cancel()
	}

	mongodbClient, err := portsrepo.NewMongoClient(mongodbAddress)
	if err != nil {
		log.Error(err)
		cancel()
	}
	defer mongodbClient.Client.Disconnect(ctx)
	grpcServer := grpc.NewPortsProtocolServer(port, mongodbClient, cancel)
	defer grpcServer.GracefulStop()
	go func() {
		if err = grpcServer.Serve(); err != nil {
			log.Error(err)
			cancel()
		}
	}()
	<-ctx.Done()
	log.Info("Stopping Port Domain Service")
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
