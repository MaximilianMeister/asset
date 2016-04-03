package asset

import (
	"fmt"
)

type higherLowerError struct {
	higher, lower float64
}

func (e *higherLowerError) Error() string {
	return fmt.Sprintf("%f - cannot be lower than %f", e.higher, e.lower)
}
