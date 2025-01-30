package model

import (
	"context"

	"github.com/samvel333/gomongo/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type Item struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}

// GetItems fetches all items from the database
func GetItems() ([]Item, error) {
	var items []Item
	collection := mongodb.GetCollection("items")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var item Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// CreateItem inserts an item into the MongoDB database
func CreateItem(item Item) error {
	collection := mongodb.GetCollection("items")
	_, err := collection.InsertOne(context.TODO(), item)
	return err
}
