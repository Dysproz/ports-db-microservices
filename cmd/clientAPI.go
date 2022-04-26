package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dysproz/ports-db-microservices/internal/core/services/grpcclient"
	jsonparser "github.com/Dysproz/ports-db-microservices/internal/core/services/jsonparse"
	"github.com/Dysproz/ports-db-microservices/internal/core/services/portsprotocol"
	"github.com/Dysproz/ports-db-microservices/internal/handlers"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	defer func() {
		log.Info("ClientAPI fully stopped")
		os.Exit(0)
	}()

	domainServerPort, serverAddress, err := getParameters()
	serverAddr := fmt.Sprintf("%v:%d", serverAddress, domainServerPort)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-sigCh
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	log.Info("Dialing ", serverAddr, "...")
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := portsprotocol.NewPortServiceClient(conn)
	stream := jsonparser.NewStream()
	grpcClient := grpcclient.NewGrpcClient(client)

	go watchJSONStream(ctx, stream, grpcClient)
	go handlers.NewRESTClient(client, stream).HandleRequests()
	<-ctx.Done()
	log.Info("Stopping JSON stream")
}

func watchJSONStream(ctx context.Context, stream *jsonparser.Stream, grpcClient *grpcclient.GrpcClient) error {
	for data := range stream.Watch() {
		if data.Error != nil {
			log.Fatal(data.Error)
			return data.Error
		}
		log.Info("creating ", data.Key, " : ", data.Port.Name)
		if err := grpcClient.CreateOrUpdatePort(data.Key, data.Port); err != nil {
			log.Fatal(err)
			return err
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
