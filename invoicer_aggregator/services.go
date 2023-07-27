package main

import (
	"fmt"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

const basePrice = 0.12

type InvoiceAggregator struct {
	data map[int]float64
}

func NewInvoiceAggregator() *InvoiceAggregator {
	return &InvoiceAggregator{
		data: make(map[int]float64),
	}
}

func (i *InvoiceAggregator) AggregateDistance(d types.CalculatedDistance) error {
	i.data[d.OBUID] += d.Distance
	return nil
}

func (i *InvoiceAggregator) Get(obuID int) (types.Invoice, error) {
	dist, ok := i.data[obuID]
	if !ok {
		return types.Invoice{}, fmt.Errorf("no information available for this OBU ID")
	}

	inv := types.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		TotalAmount:   dist * basePrice,
	}
	return inv, nil
}
