package compute

import (
	"errors"
	"strings"

	"go.uber.org/zap"
)

var (
	errInvalidQuery     = errors.New("empty query")
	errInvalidCommand   = errors.New("invalid command")
	errInvalidArguments = errors.New("invalid arguments amount")
)

type Compute struct {
	Logger *zap.Logger
}
type Query struct {
	Code      int
	Arguments []string
}

var commands = map[string]int{
	"SET": 1,
	"GET": 2,
	"DEL": 3,
}

var needArguments = map[int]int{
	1: 2,
	2: 1,
	3: 1,
}

func (d *Compute) Parse(query []byte) (Query, error) {

	tokens := strings.Fields(string(query))

	if len(tokens) == 0 {
		d.Logger.Debug("No Commands were endered")
		return Query{}, errInvalidQuery
	}

	if _, ok := commands[tokens[0]]; !ok {
		d.Logger.Debug("Wrong Command was entered")
		return Query{}, errInvalidCommand
	}

	code := commands[tokens[0]]

	if len(tokens) < 2 {
		d.Logger.Debug("No arguments were provided")
		return Query{}, errInvalidArguments
	}

	if needArguments[code] != len(tokens)-1 {
		return Query{}, errInvalidArguments
	}

	var arguments []string
	for _, token := range tokens[1:] {
		arguments = append(arguments, token)
	}

	return Query{
		Code:      commands[tokens[0]],
		Arguments: arguments}, nil

}
