package main

import (
	"context"
	"fmt"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	conn *mongo.Client
	col  *mongo.Collection
}

func NewMongoStore() *MongoStore {
	conn, err := mongo.Connect(context.Background())
	if err != nil {
		fmt.Println("mongo error:", err.Error())
		return nil
	}

	col := conn.Database("Toll-Calculator").Collection("Trips")

	return &MongoStore{
		conn: conn,
		col:  col,
	}
}

func (s *MongoStore) InsertOBUDataInDB(ctx context.Context, data types.CalculatedDistance) error {
	_, err := s.col.InsertOne(ctx, data)
	if err != nil {
		fmt.Println("mongo error:", err.Error())
	}
	return nil
}

func (s *MongoStore) GetInvoice(ctx context.Context, obuID int64, date string) (float64, error) {
	filter := bson.M{"obuid": obuID, "date": date}

	cursor, err := s.col.Find(ctx, filter)
	if err != nil {
		return 0, err
	}

	var sum float64
	for cursor.Next(ctx) {
		var data types.CalculatedDistance
		if err := cursor.Decode(&data); err != nil {
			return 0, err
		}
		sum = sum + data.Distance
	}

	if err := cursor.Err(); err != nil {
		return 0, err
	}

	return sum, nil

}
