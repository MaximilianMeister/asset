package asset

import (
	"github.com/shopspring/decimal"
)

var stopLossTests = []struct {
	actual      float64
	stop        float64
	expected    float64
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
	{
		"consors",
		Broker{
			"Consors Bank",
			decimal.NewFromFloat(4.95),
			decimal.NewFromFloat(0.0025),
			decimal.NewFromFloat(9.95),
			decimal.NewFromFloat(69.0),
		},
	},
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
	volume      int64
	target      float64
	actual      float64
	stop        float64
	commission  float64
	amount      int64
	gain        float64
	loss        float64
	even        float64
	rrr         decimal.Decimal
}{
	{"consors", 1000, 3.56, 3.01, 2.87, 19.9, 332, 162.02, 67.06, 3.07, decimal.NewFromFloat(3.9)},
	{"ingdiba", 2000, 38.56, 32.01, 29.87, 19.8, 62, 370.92, 167.86, 32.58, decimal.NewFromFloat(3.1)},
	{"comdirect", 29000, 385.06, 327.01, 298.87, 119.8, 89, 5150.54, 2520.37, 327.19, decimal.NewFromFloat(2.1)},
}
