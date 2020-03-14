package repository

import (
	"context"
	Data "github.com/jdschrack/mongotutorial/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var tableName string = "people"

func AddPerson(person Data.Person) (*mongo.InsertOneResult, error) {
	log.Printf("Creating %s", person)
	collection := Data.Database.Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := collection.InsertOne(ctx, person)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllPeople() ([]Data.Person, error) {
	log.Println("GetPerson all People")
	var people []Data.Person
	collect := Data.Database.Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := collect.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person Data.Person
		cursor.Decode(&person)
		people = append(people, person)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return people, nil
}

func GetPerson(id primitive.ObjectID) (*Data.Person, error) {
	log.Println(id)
	var person Data.Person
	collection := Data.Database.Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := collection.FindOne(ctx, Data.Person{ID: id}).Decode(&person)

	if err != nil {
		return nil, err
	}

	return &person, nil
}
