package initialization

import (
	"log"
	"spider/internal/database/compute"

	"go.uber.org/zap"
)

type Initializer struct {
	Server Server
}

func CreateInitializer() (Initializer, error) {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	compute := compute.CreateCompute(logger)

	server, err := CreateServer(compute)
	if err != nil {
		log.Println(err)
		return Initializer{}, err
	}

	initializer := Initializer{
		Server: server,
	}

	return initializer, nil

}
