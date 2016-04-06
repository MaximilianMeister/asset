package asset

import (
	"fmt"
)

type HigherLowerError struct {
	higher, lower float64
}

func (e *HigherLowerError) Error() string {
	return fmt.Sprintf("%f - cannot be lower than %f", e.higher, e.lower)
}

type InvalidBrokerError struct {
	brokerAlias string
}

func (e *InvalidBrokerError) Error() string {
	return fmt.Sprintf("%s is not a valid broker", e.brokerAlias)
}
