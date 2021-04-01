package repository

import (
	"context"
	"log"
	"time"

	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/database"
	"github.com/KinyaElGrande/Go-exercise-101/graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dogsCollection *mongo.Collection = database.OpenCollection(database.Client, "dogs")

func Save(input *model.NewDog) *model.Dog {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := dogsCollection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Dog{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Name:      input.Name,
		IsGoodBoy: input.IsGoodBoy,
	}
}

func FindByID(ID string) *model.Dog {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := dogsCollection.FindOne(ctx, bson.M{"_id": objectId})
	dog := model.Dog{}
	res.Decode(&dog)
	return &dog
}

func AllDogs() []*model.Dog {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := dogsCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var dogs []*model.Dog

	for cur.Next(ctx) {
		var dog *model.Dog
		err := cur.Decode(&dog)
		if err != nil {
			log.Fatal(err)
		}
		dogs = append(dogs, dog)
	}

	return dogs
}
