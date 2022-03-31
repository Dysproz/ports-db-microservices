package mongodb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Dysproz/ports-db-microservices/pkg/portsprotocol"
)

// PortDB reflects structure kept in mongoDB database
type PortDB struct {
	Key  string  `bson:"key"`
	Port pb.Port `bson:"port"`
}

// MongoClient is handling communiction with mongoD database
type MongoClient struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// InsertOrUpdatePort checks if key already exists and replaces entry or inserts a new one
func (m *MongoClient) InsertOrUpdatePort(key string, port pb.Port) error {
	filter := bson.D{{"key", key}}
	portEntry := PortDB{
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

// GetPort gets from database port data by passed key
func (m *MongoClient) GetPort(key string) (pb.Port, error) {
	var port PortDB
	filter := bson.D{{"key", key}}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := m.Collection.FindOne(ctx, filter).Decode(&port); err != nil {
		log.Debug("Get port for key: ", key, " failed with error: ", err)
		return pb.Port{}, err
	}
	log.Debug("Found port for key: ", key, " with data: ", port)
	return port.Port, nil
}

// Close disconnects from mongoDB database
func (m *MongoClient) Close() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Info("Connection to MongoDB closed.")
}

// CreateMongoDBClient creates a new connection to mongoDB database
func CreateMongoDBClient(mongodbAddr string) (MongoClient, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v", mongodbAddr))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return MongoClient{}, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return MongoClient{}, err
	}

	log.Info("Connected to MongoDB at ", mongodbAddr)
	collection := client.Database("port").Collection("port")
	return MongoClient{
		Client:     client,
		Collection: collection,
	}, nil
}
