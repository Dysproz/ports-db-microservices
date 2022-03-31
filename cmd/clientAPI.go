package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	clientgrpc "github.com/Dysproz/ports-db-microservices/pkg/grpc"
	"github.com/Dysproz/ports-db-microservices/pkg/jsonparser"
	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
	"github.com/Dysproz/ports-db-microservices/pkg/resthandler"
)

func main() {
	defer os.Exit(0)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	domainServerPort, serverAddress, err := getParameters()
	serverAddr := fmt.Sprintf("%v:%d", serverAddress, domainServerPort)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Info("Dialing ", serverAddr, "...")
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPortServiceClient(conn)
	stream := jsonparser.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Info(data.Error)
			}
			log.Info("creating ", data.Key, " : ", data.Port.Name)
			if err := clientgrpc.CreateOrUpdatePort(client, data.Key, data.Port); err != nil {
				log.Fatal(err)
			}
		}
	}()
	go resthandler.HandleRequests(client, stream)
	select {
	case <-sigCh:
		log.Info("Interrupt signal detected. Gracefully shutting down...")
		runtime.Goexit()
	}
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
