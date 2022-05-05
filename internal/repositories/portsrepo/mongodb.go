package portsrepo

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Dysproz/ports-db-microservices/internal/core/domain"
)

type portDB struct {
	Key  string      `bson:"key"`
	Port domain.Port `bson:"port"`
}

// MongoClient is a mongoDB client
type MongoClient struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// NewMongoClient returns a mongoDB client with connection opened at specified address
func NewMongoClient(address string) (*MongoClient, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v", address))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return &MongoClient{}, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return &MongoClient{}, err
	}
	log.Info("Connected to MongoDB at ", address)
	collection := client.Database("port").Collection("port")

	return &MongoClient{
		Client:     client,
		Collection: collection,
	}, nil
}

// InsertOrUpdate checks if key already exists and replaces entry or inserts a new one
func (m *MongoClient) InsertOrUpdate(key string, port domain.Port) error {
	filter := bson.D{{"key", key}}
	portEntry := portDB{
		Key:  key,
		Port: port,
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := m.Collection.FindOne(ctx, filter).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			insertResult, err := m.Collection.InsertOne(context.TODO(), portEntry)
			log.Debug("Insert operation for ", key, " resulted with: ", insertResult)
			return err
		}
	}
	replaceResult, err := m.Collection.ReplaceOne(context.TODO(), filter, portEntry)
	log.Debug("Replace operation for ", key, " resulted with: ", replaceResult)
	return err
}

// Get gets from database port data by passed key
func (m *MongoClient) Get(key string) (domain.Port, error) {
	var port portDB
	filter := bson.D{{"key", key}}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := m.Collection.FindOne(ctx, filter).Decode(&port); err != nil {
		log.Debug("Get port for key: ", key, " failed with error: ", err)
		return domain.Port{}, err
	}
	log.Debug("Found port for key: ", key, " with data: ", port)
	return port.Port, nil
}
