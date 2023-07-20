package main

import (
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

// This will handle invoice requests and services.
type Aggregator interface {
	AggregateDistance(types.OBUData) error
}

type Storer interface {
	Insert(types.CalculatedDistance) error
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) *InvoiceAggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.CalculatedDistance) error {
	return i.store.Insert(distance)
}
