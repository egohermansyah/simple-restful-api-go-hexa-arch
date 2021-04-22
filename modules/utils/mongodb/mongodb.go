package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func Connect(uri string, dbName string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)

	return database, nil
}

func NewIndexing(collection *mongo.Collection, options *options.IndexOptions, data ...[]string) error {

	for _, values := range data {

		document := bsonx.Doc{}

		for _, value := range values {
			document = append(document, bsonx.Elem{Key: value, Value: bsonx.Int32(1)})
		}

		err := InitIndex(collection, options, document)
		if err != nil {
			return err
		}
	}

	return nil
}

func InitIndex(collection *mongo.Collection, options *options.IndexOptions, data bsonx.Doc) error {
	var err error

	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    data,
			Options: options,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
