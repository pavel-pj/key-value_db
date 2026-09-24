package initialization

import (
	"errors"
	"spider/internal/database/compute"

	"go.uber.org/zap"
)

var errorLoggerEmpty = errors.New("logger structer is epmty")

func CreateCompute(logger *zap.Logger) (compute.Compute, error) {
	if logger == nil {

		return compute.Compute{}, errorLoggerEmpty
	}
	compute := compute.Compute{
		Logger: logger,
	}
	return compute, nil

}
