package broker

import (
	"github.com/shopspring/decimal"
)

var newBrokersTests = map[string]Broker{
	"onvista": {
		"OnVista Bank",
		decimal.NewFromFloat(5.99),
		decimal.NewFromFloat(0.0023),
		decimal.NewFromFloat(5.99),
		decimal.NewFromFloat(39),
	},
	"dab": {
		"DAB Bank",
		decimal.NewFromFloat(4.95),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(3.99),
		decimal.NewFromFloat(55),
	},
	"targo": {
		"Targo Bank",
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(8.9),
		decimal.NewFromFloat(34.9),
	},
	"consors": {
		"Consors Bank",
		decimal.NewFromFloat(4.95),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.95),
		decimal.NewFromFloat(69),
	},
	"ingdiba": {
		"ING Diba",
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.9),
		decimal.NewFromFloat(59.9),
	},
	"comdirect": {
		".comdirect",
		decimal.NewFromFloat(4.9),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.9),
		decimal.NewFromFloat(59.9),
	},
	"sbroker": {
		"SBroker",
		decimal.NewFromFloat(4.95),
		decimal.NewFromFloat(0.0025),
		decimal.NewFromFloat(9.95),
		decimal.NewFromFloat(49.95),
	},
	"maxblue": {
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
