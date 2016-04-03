package asset

var stopLossTests = []struct {
	actual   float64
	stop     float64
	expected float64
	errExpected bool
}{
	{3.4, 3.0, 3.0, false},
	{3.2, 3.2, 3.2, true}, // stop loss same as actual price
	{3.2, 4.5, 3.2, true}, // stop loss higher than actual price
}

var findBrokerTests = []struct {
	brokerAlias string
	expected    Broker
}{
	{"consors", Broker{"Consors Bank", 4.95, 0.0025, 9.95, 69.0}},
}

var isBrokerTests = []struct {
	brokerAlias string
	expected    bool
}{
	{"consors", true},
	{"bonsors", false},
}

var orderTests = []struct {
	brokerAlias string
	volume      uint32
	target      float64
	actual      float64
	stop        float64
	commission  float64
	amount      uint32
	gain        float64
	loss        float64
	even        float64
	rrr         float64
}{
	{"consors", 1000, 3.56, 3.01, 2.87, 19.9, 332, 162.03, 67.06, 3.08, 2.1},
	{"ingdiba", 2000, 38.56, 32.01, 29.87, 19.8, 62, 370.93, 167.86, 32.58, 3.1},
	{"comdirect", 29000, 385.06, 327.01, 298.87, 119.8, 88, 4765.48, 2819.24, 330.91, 3.9},
}
