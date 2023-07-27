package main

import (
	"context"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type InvoiceAggregator struct {
	data       map[int]float64
	mongoStore *MongoStore
}

func NewInvoiceAggregator() *InvoiceAggregator {
	mongoStore := NewMongoStore()

	return &InvoiceAggregator{
		data:       make(map[int]float64),
		mongoStore: mongoStore,
	}
}

func (i *InvoiceAggregator) AggregateDistance(d types.CalculatedDistance) error {
	i.data[d.OBUID] += d.Distance
	i.mongoStore.InsertOBUDataInDB(context.Background(), d)
	return nil
}

func (i *InvoiceAggregator) GetInvoice(obuID int) ([]*types.Invoice, error) {
	resp, error := i.mongoStore.GetInvoice(context.Background(), int64(obuID))

	if error != nil {
		return nil, error
	}
	return resp, nil
}
