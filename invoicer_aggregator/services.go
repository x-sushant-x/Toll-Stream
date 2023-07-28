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

func (i *InvoiceAggregator) GetInvoice(obuID int, date string) (float64, error) {
	resp, error := i.mongoStore.GetInvoice(context.Background(), int64(obuID), date)

	if error != nil {
		return 0, error
	}
	return resp, nil
}
