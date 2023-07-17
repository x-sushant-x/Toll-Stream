package main

import "github.com/sushant102004/Traffic-Toll-Microservice/types"

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Insert(d types.CalculatedDistance) error {
	return nil
}
