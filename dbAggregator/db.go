package main

import (
	"context"
	"fmt"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
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
