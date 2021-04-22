package role

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"simple-restful-api-go-hexa-arch/business/role/defined"
	"time"
)

type Collection struct {
	Id       string    `bson:"_id"`
	Name     string    `bson:"name"`
	Desc     string    `bson:"desc"`
	Created  time.Time `bson:"created"`
	Modified time.Time `bson:"modified"`
}

func NewCollection(role defined.Role) *Collection {
	return &Collection{
		Id:   role.Id,
		Name: role.Name,
		Desc: role.Desc,
	}
}

func (collection *Collection) MappingData() defined.Role {
	var data defined.Role
	data.Id = collection.Id
	data.Name = collection.Name
	data.Desc = collection.Desc
	data.Created = collection.Created
	data.Modified = collection.Modified
	return data
}

type MongoDBRepository struct {
	collection *mongo.Collection
}

func NewMongoDBRepository(db *mongo.Database) (*MongoDBRepository, error) {
	repository := MongoDBRepository{db.Collection("role")}
	return &repository, nil
}

func (repository *MongoDBRepository) Insert(role defined.Role) (*defined.Role, error) {
	collection := NewCollection(role)
	id := primitive.NewObjectID().Hex()
	collection.Id = id
	collection.Created = time.Now()
	collection.Modified = time.Now()
	_, err := repository.collection.InsertOne(context.TODO(), collection)
	if err != nil {
		return nil, err
	}
	result := collection.MappingData()
	return &result, nil
}

func (repository *MongoDBRepository) FindById(id string) (*defined.Role, error) {
	var collection Collection
	filter := bson.M{
		"_id": id,
	}
	err := repository.collection.FindOne(context.TODO(), filter).Decode(&collection)
	if err != nil {
		return nil, err
	}
	result := collection.MappingData()
	result.Id = id
	return &result, nil
}

func (repository *MongoDBRepository) UpdateById(id string, name string, desc string) (*defined.Role, error) {
	data := bson.M{
		"$set": bson.M{
			"name": name,
			"desc": desc,
		},
	}
	_, err := repository.collection.UpdateByID(context.TODO(), id, data)
	if err != nil {
		return nil, err
	}
	result, err := repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *MongoDBRepository) DeleteById(id string) error {
	filter := bson.M{
		"_id": id,
	}
	result, err := repository.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("mongo: no documents in result")
	}
	return nil
}

func (repository *MongoDBRepository) List(skip int, perPage int) ([]defined.Role, error) {
	var roles []defined.Role
	option := options.Find()
	option.SetSort(bson.D{{Key: "_id", Value: 1}})
	option.SetSkip(int64(skip))
	option.SetLimit(int64(perPage))
	cursor, err := repository.collection.Find(context.TODO(), bson.M{}, option)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(context.TODO()); err != nil {
			//todo - error log for close mongodb cursor
		}
	}()
	for cursor.Next(context.TODO()) {
		var collection Collection
		err := cursor.Decode(&collection)
		if err != nil {
			return nil, err
		}
		data := collection.MappingData()
		roles = append(roles, data)
	}
	return roles, nil
}
