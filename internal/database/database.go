package database

import (
	"errors"
	"spider/internal/database/compute"
	"spider/internal/database/storage/engine"
)

type Database struct {
	Compute compute.Compute
	Engine  *engine.HashTable
}

func CreateDatabase(compute compute.Compute, engine engine.HashTable) (*Database, error) {

	if compute.Logger == nil {
		return nil, errors.New("compute is invalid")
	}

	//if engine.Data == nil {
	//	return nil, errors.New("storage is invalid")
	//}

	return &Database{
		Compute: compute,
		Engine:  &engine,
	}, nil
}

func (d *Database) HandleQueries(query []byte) (string, error) {

	data, err := d.Compute.Parse(query)

	if err != nil {
		return "", err
	}

	switch data.Code {
	case 1:
		return d.handlerSET(data)
	case 2:
		return d.handlerGET(data)
	case 3:
		return d.handlerDEL(data)
	}

	return "", errors.New("No handler in database")
	//return fmt.Sprintf("[ok] %d", data.Code)

}

func (d *Database) handlerSET(query compute.Query) (string, error) {
	d.Engine.Set(query)
	return "Ok", nil

}

func (d *Database) handlerGET(query compute.Query) (string, error) {
	result, err := d.Engine.GET(query)
	if err != nil {
		return "", err
	}
	return result, nil

}

func (d *Database) handlerDEL(query compute.Query) (string, error) {
	result, err := d.Engine.DEL(query)
	if err != nil {
		return "", err
	}
	return result, nil

}
