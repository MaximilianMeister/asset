package asset

var roundUpTests = []struct {
	input  float64
	digits int
	output float64
}{
	{10.5678, 2, 10.57},
}

var roundDownTests = []struct {
	input  float64
	digits int
	output float64
}{
	{10.5678, 2, 10.56},
}
