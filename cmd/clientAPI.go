package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/Dysproz/ports-db-microservices/internal/datainput"
	"github.com/Dysproz/ports-db-microservices/internal/grpc"
	"github.com/Dysproz/ports-db-microservices/internal/handlers"
)

func main() {
	defer func() {
		log.Info("ClientAPI fully stopped")
	}()
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-sigCh
		log.Error("system call:%+v", oscall)
		cancel()
	}()

	domainServerPort, serverAddress, err := getParameters()
	if err != nil {
		log.Error(err)
		cancel()
	}
	serverAddr := fmt.Sprintf("%v:%d", serverAddress, domainServerPort)

	stream := datainput.NewStream()
	Client := grpc.NewClient(serverAddr, cancel)
	defer Client.CloseConnection()

	go watchJSONStream(cancel, stream, Client)
	go handlers.NewRESTClient(Client, stream).HandleRequests(cancel)
	<-ctx.Done()
	log.Info("Stopping JSON stream")
}

func watchJSONStream(cancel context.CancelFunc, stream *datainput.Stream, Client *grpc.Client) error {
	for data := range stream.Watch() {
		if data.Error != nil {
			log.Error(data.Error)
			cancel()
		}
		log.Info("creating ", data.Key, " : ", data.Port.Name)
		if err := Client.CreateOrUpdatePort(data.Key, data.Port); err != nil {
			log.Error(err)
			cancel()
		}
	}
	return nil
}

func getParameters() (int, string, error) {
	domainServerPort := flag.Int("domain-server-port", 5000, "Domain server port.")
	serverAddress := flag.String("server-address", "portdomainserver_1", "Address of the port domain server.")
	logLevel := flag.String("log-level", "info", "Log level.")
	flag.Parse()
	logLevelParsed, err := log.ParseLevel(*logLevel)
	if err != nil {
		return 0, "", err
	}
	log.SetLevel(logLevelParsed)
	log.Debug("Running with CLI parameters. domain-server-port: ", *domainServerPort,
		" server-address: ", *serverAddress,
		" log-level: ", *logLevel)
	return *domainServerPort, *serverAddress, nil
}
