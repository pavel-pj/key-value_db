package engine

import (
	"errors"
	"spider/internal/database/compute"
)

var errNotFound = errors.New("key not found")

type HashTable struct {
	Data map[string]string
}

func CreateEngine() *HashTable {

	d := make(map[string]string)

	return &HashTable{
		Data: d,
	}
}

func (e *HashTable) Set(query compute.Query) error {

	e.Data[query.Arguments[0]] = query.Arguments[1]

	return nil

}

func (e *HashTable) GET(query compute.Query) (string, error) {

	if _, ok := e.Data[query.Arguments[0]]; !ok {

		return "", errNotFound
	}

	return e.Data[query.Arguments[0]], nil

}

func (e *HashTable) DEL(query compute.Query) (string, error) {

	if _, ok := e.Data[query.Arguments[0]]; !ok {

		return "", errNotFound
	}

	delete(e.Data, query.Arguments[0])

	return "OK", nil

}
