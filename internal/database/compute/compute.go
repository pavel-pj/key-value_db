package compute

import (
	"errors"
	"log"
	"strings"

	"go.uber.org/zap"
)

var (
	errInvalidQuery     = errors.New("empty query")
	errInvalidCommand   = errors.New("invalid command")
	errInvalidArguments = errors.New("invalid arguments")
)

type Compute struct {
	logger *zap.Logger
}
type Query struct {
	Command   string
	Arguments []string
}

var commands = map[string]int{
	"SET": 1,
	"GET": 2,
	"DEL": 3,
}

func CreateCompute(logger *zap.Logger) Compute {
	if logger == nil {
		log.Fatal("logger structer is epmty")

	}

	compute := Compute{
		logger: logger,
	}
	return compute

}

func (d *Compute) Parse(query string) (Query, error) {

	tokens := strings.Fields(query)

	if len(tokens) == 0 {
		d.logger.Debug("No Commands were endered")
		return Query{}, errInvalidQuery
	}

	if _, ok := commands[tokens[0]]; !ok {
		d.logger.Debug("Wrong Command was entered")
		return Query{}, errInvalidCommand
	}

	if len(tokens) < 2 {
		d.logger.Debug("No arguments were provided")
		return Query{}, errInvalidArguments
	}

	var arguments []string
	for _, token := range tokens[1:] {
		arguments = append(arguments, token)
	}

	return Query{
		Command:   tokens[0],
		Arguments: arguments}, nil

}
