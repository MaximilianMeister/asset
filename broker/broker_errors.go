package broker

import (
	"fmt"
)

// InvalidBrokerError represents an error for a non existing Broker
type InvalidBrokerError struct {
	brokerAlias string
}

func (e *InvalidBrokerError) Error() string {
	return fmt.Sprintf("%s is not a valid broker", e.brokerAlias)
}
