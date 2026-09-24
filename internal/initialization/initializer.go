package initialization

import (
	"errors"
	"log"
	"spider/internal/database"
	"spider/internal/database/storage/engine"
	"spider/internal/network"

	"go.uber.org/zap"
)

type Initializer struct {
	Engine *engine.HashTable
	Logger *zap.Logger
	Server network.Server
}

func CreateInitializer() (Initializer, error) {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server, err := network.CreateServer()
	if err != nil {
		log.Println(err)
		return Initializer{}, err
	}

	engine := engine.CreateEngine()

	initializer := Initializer{
		Engine: engine,
		Logger: logger,
		Server: server,
	}

	return initializer, nil

}

func StartDatabase(i Initializer) error {

	compute, err := CreateCompute(i.Logger)

	if err != nil {
		log.Println(err)
		return errors.New("Empty Compute layer")
	}

	db, err := database.CreateDatabase(compute, *i.Engine)
	if err != nil {
		log.Println(err)
		return errors.New("Empty database")
	}

	i.Server.HandleQueries(func(query []byte) []byte {

		response, err := db.HandleQueries(query)
		if err != nil {
			return []byte(err.Error())
		}
		return []byte(response)
	})

	return nil

}
