// Package asset provides an api to calculate asset price projections.
package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/shopspring/decimal"
)

// static broker data
type Broker struct {
	Name           string          `json:"name"`
	BasicPrice     decimal.Decimal `json:"basic_price"`
	CommissionRate decimal.Decimal `json:"commission_rate"`
	MinRate        decimal.Decimal `json:"min_rate"`
	MaxRate        decimal.Decimal `json:"max_rate"`
}

// map of all brokers and their static data
type Brokers map[string]Broker

// data to calculate order figures
type Order struct {
	brokerAlias          string
	volume               int64
	target, actual, stop decimal.Decimal
}

// returns a map of type Brokers which contains all static broker data
// defined in brokers.json
func BrokerRegister() (brokers Brokers, err error) {
	// please note that the static data can be outdated and does not contain all
	// brokers.
	// if your broker and it's rates are missing please add it, and don't hesitate
	// to send a pull request to https://github.com/MaximilianMeister/asset
	file, err := ioutil.ReadFile("./broker.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	if err = json.Unmarshal(file, &brokers); err != nil {
		return brokers, err
	}

	return brokers, nil
}

// IsBroker returns bool if broker is available
func IsBroker(brokerAlias string) (bool, error) {
	register, err := BrokerRegister()
	if err != nil {
		return false, err
	}

	for b := range register {
		if b == brokerAlias {
			return true, nil
		}
	}

	return false, nil
}

// FindBroker returns a Broker type with all static data about a single broker
func FindBroker(brokerAlias string) (Broker, error) {
	valid, err := IsBroker(brokerAlias)
	if err != nil {
		return Broker{}, err
	}

	if valid {
		register, err := BrokerRegister()
		if err != nil {
			return Broker{}, err
		}

		for b := range register {
			if b == brokerAlias {
				return register[b], nil
			}
		}
	}

	return Broker{}, nil
}

// StopLoss returns a stop loss value for an order when it is valid.
// when invalid it returns an error along with the actual price
func StopLoss(actual, stop float64) (float64, error) {
	if stop >= actual {
		return actual, &higherLowerError{stop, actual}
	}

	return stop, nil
}

// RiskRewwardRatio returns a risk reward ratio value for an order
func (o *Order) RiskRewardRatio() (rrr decimal.Decimal) {
	chance := o.target.Sub(o.actual)
	risk := o.actual.Sub(o.stop)
	rrr = chance.Div(risk).Round(1)

	return
}

// TotalCommission returns the total broker commission fee for an order
func (o *Order) TotalCommission(brokerAlias string) (commission decimal.Decimal, err error) {
	commission = decimal.NewFromFloat(0.0)

	broker, err := FindBroker(brokerAlias)
	if err != nil {
		return commission, err
	}

	volumeRateBuy := decimal.New(o.Amount(), 0).Mul(o.actual).Mul(broker.CommissionRate)
	volumeRateSell := decimal.New(o.Amount(), 0).Mul(o.target).Mul(broker.CommissionRate)

	buySell := []decimal.Decimal{
		volumeRateBuy,
		volumeRateSell,
	}

	for _, bs := range buySell {
		if (bs.Add(broker.BasicPrice)).Cmp(broker.MinRate) == 1 {
			if (bs.Add(broker.BasicPrice)).Cmp(broker.MaxRate) == 1 {
				commission = commission.Add(broker.MaxRate)
			} else {
				commission = commission.Add(broker.BasicPrice.Add(bs))
			}
		} else {
			commission = commission.Add(broker.MinRate)
		}
	}
	commission = commission.Round(2)

	return
}

// Amount returns the actual amount of stocks that can be bought for an order
func (o *Order) Amount() (amount int64) {
	amount = decimal.New(o.volume, 0).Div(o.actual).Round(0).IntPart()

	return
}

// Gain returns the highest possible gain for an order
func (o *Order) Gain(broker string) (gain decimal.Decimal, err error) {
	amount := decimal.New(o.Amount(), 0)
	volume := decimal.New(o.volume, 0)
	commission, err := o.TotalCommission(broker)
	if err != nil {
		return decimal.NewFromFloat(0.0), err
	}

	gain = amount.Mul(o.target)
	gain = gain.Sub(volume).Sub(commission)
	gain = gain.Round(2)

	return gain, nil
}

// Loss returns the highest possible loss for an order
func (o *Order) Loss(broker string) (loss decimal.Decimal, err error) {
	amount := decimal.New(o.Amount(), 0)
	volume := decimal.New(o.volume, 0)
	commission, err := o.TotalCommission(broker)
	if err != nil {
		return decimal.NewFromFloat(0.0), err
	}

	loss = volume.Sub(amount.Mul(o.stop)).Add(commission)
	loss = loss.Round(2)

	return loss, nil
}

// Even returns the exact break even point for an order
func (o *Order) Even(broker string) (even decimal.Decimal, err error) {
	amount := decimal.New(o.Amount(), 0)
	volume := decimal.New(o.volume, 0)
	commission, err := o.TotalCommission(broker)
	if err != nil {
		return decimal.NewFromFloat(0.0), err
	}

	even = volume.Add(commission)
	even = even.Div(amount)
	even = even.Round(2)

	return even, nil
}
