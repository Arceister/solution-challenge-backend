package repository

import (
	"api-solution/lib"
	"api-solution/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type PreferableRepository struct {
	mongo lib.Mongo
}

func NewPreferableRepository(mongo lib.Mongo) PreferableRepository {
	return PreferableRepository{
		mongo: mongo,
	}
}

func (r PreferableRepository) GetAll() (preferables []models.UserPreferable, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	query, err := r.mongo.Mongo.Collection("preferables").Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("error", err)
		return preferables, err
	}
	defer query.Close(ctx)

	listPreferables := make([]models.UserPreferable, 0)
	for query.Next(ctx) {
		var row models.UserPreferable
		err := query.Decode(&row)
		if err != nil {
			fmt.Println("Error")
		}
		listPreferables = append(listPreferables, row)
	}

	preferables = listPreferables
	return preferables, err
}

func (r PreferableRepository) InsertOne(insertable models.UserPreferable) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.mongo.Mongo.Collection("preferables").InsertOne(ctx, insertable)
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	return nil
}
