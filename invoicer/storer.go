/*
	Purpose of this file:
		Provide Insert method that will add distance of an OBU into a map.
*/

package main

import (
	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStore) Insert(d types.CalculatedDistance) error {
	m.data[d.OBUID] += d.Distance
	return nil
}
