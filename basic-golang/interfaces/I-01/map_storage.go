package main

import (
	"errors"
)

type MapStorage struct {
	storage map[int]int
}

func newMapStorage() *MapStorage {
	return &MapStorage{storage: make(map[int]int)}
}

func (m *MapStorage) Set(key int, value int) error {
	m.storage[key] = value
	return nil
}

func (m *MapStorage) Get(k int) (int, error) {
	if value, ok := m.storage[k]; ok {
		return value, nil
	}
	return -1, errors.New("Not Exists")
}

func (m *MapStorage) Check(k int) bool {
	_, exists := m.storage[k]
	return exists
}

func (m *MapStorage) Remove(k int) error {
	delete(m.storage, k)
	return nil
}
