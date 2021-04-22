package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-restful-api-go-hexa-arch/config"
	"simple-restful-api-go-hexa-arch/modules/utils/mongodb"
)

func ConnectDatabase(dbConfig config.MongoDbConfig) *mongo.Database {
	uri := fmt.Sprintf("%s://", dbConfig.Driver)
	if dbConfig.Port == 0 {
		uri = fmt.Sprintf("%s+srv://", dbConfig.Driver)
	}
	uri = fmt.Sprintf("%s%v:%v/?connect=direct",
		uri,
		dbConfig.Host,
		dbConfig.Port)
	db, err := mongodb.Connect(uri, dbConfig.Name)
	if err != nil {
		panic(err)
	}
	return db
}
