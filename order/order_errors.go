package order

import (
	"fmt"
)

// HigherLowerError represents an error for values (a & b) where a cannot be lower than b
type HigherLowerError struct {
	higher, lower float64
}

func (e *HigherLowerError) Error() string {
	return fmt.Sprintf("%f - cannot be lower than %f", e.higher, e.lower)
}
