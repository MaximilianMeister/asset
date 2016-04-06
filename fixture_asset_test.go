package asset

import (
	"github.com/shopspring/decimal"
)

var newBrokersTests = map[string]Broker{
	"onvista": Broker{
		"OnVista Bank",
		decimal.NewFromFloat(5.99),
		decimal.NewFromFloat(0.0023),
		decimal.NewFromFloat(5.99),
		decimal.NewFromFloat(39),
	},
	"dab": Broker{
		"DAB Bank",
		decimal.NewFromFloat(4.95),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(3.99),
		decimal.NewFromFloat(55),
	},
	"targo": Broker{
		"Targo Bank",
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(8.9),
		decimal.NewFromFloat(34.9),
	},
	"consors": Broker{
		"Consors Bank",
		decimal.NewFromFloat(4.95),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.95),
		decimal.NewFromFloat(69),
	},
	"ingdiba": Broker{
		"ING Diba",
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.9),
		decimal.NewFromFloat(59.9),
	},
	"comdirect": Broker{
		".comdirect",
		decimal.NewFromFloat(4.9),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.9),
		decimal.NewFromFloat(59.9),
	},
	"sbroker": Broker{
		"SBroker",
		decimal.NewFromFloat(4.95),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.95),
		decimal.NewFromFloat(49.95),
	},
	"maxblue": Broker{
		"maxblue",
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(7.9),
		decimal.NewFromFloat(39.9),
	},
}

var isBrokerTests = []struct {
	brokerAlias string
	errExpected bool
}{
	{"consors", false},
	{"bonsors", true},
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
	{"dab", 5000, 38.56, 32.01, 29.87, 37.42, 156, 977.94, 377.7, 32.29, decimal.NewFromFloat(3.1)},
	{"comdirect", 29000, 385.06, 327.01, 298.87, 119.8, 89, 5150.54, 2520.37, 327.19, decimal.NewFromFloat(2.1)},
}

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
