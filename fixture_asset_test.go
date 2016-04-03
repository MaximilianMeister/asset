package asset

var stopLossTests = []struct {
	actual   float32
	stop     float32
	expected float32
}{
	{3.4, 3.0, 3.0},
	{3.2, 4.5, 3.2}, // stop loss higher than or same as actual price
}

var getBrokerTests = []struct {
	shortName string
	expected  Broker
}{
	{"consors", Broker{"Consors Bank", 4.95, 0.0025, 9.95, 69.0}},
}

var isBrokerTests = []struct {
	shortName string
	expected  bool
}{
	{"consors", true},
	{"bonsors", false},
}

var orderTests = []struct {
	broker     string
	volume     uint32
	target     float32
	actual     float32
	stop       float32
	commission float32
	amount     uint32
	gain       float32
	loss       float32
	even       float32
	rrr        float32
}{
	{"consors", 1000, 3.56, 3.01, 2.87, 19.9, 332, 162.02, 67.07, 3.08, 2.1},
	{"ingdiba", 2000, 38.56, 32.01, 29.87, 19.8, 62, 370.92, 167.86, 32.58, 3.1},
	{"comdirect", 29000, 385.06, 327.01, 298.87, 119.81, 88, 4765.48, 2819.26, 330.91, 3.9},
}
