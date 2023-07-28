package main

import (
	"context"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type InvoiceAggregator struct {
	mongoStore *MongoStore
}

func NewInvoiceAggregator() *InvoiceAggregator {
	mongoStore := NewMongoStore()

	return &InvoiceAggregator{
		mongoStore: mongoStore,
	}
}

func (i *InvoiceAggregator) AggregateDistance(d types.CalculatedDistance) error {
	i.mongoStore.InsertOBUDataInDB(context.Background(), d)
	return nil
}
