package asset

import (
	"math"
)

func RoundUp(input float64, digits int) (output float64) {
	var round float64
	pow := math.Pow(10, float64(digits))
	digit := pow * input
	round = math.Ceil(digit)
	output = round / pow
	return
}

func RoundDown(input float64, digits int) (output float64) {
	var round float64
	pow := math.Pow(10, float64(digits))
	digit := pow * input
	round = math.Floor(digit)
	output = round / pow
	return
}
