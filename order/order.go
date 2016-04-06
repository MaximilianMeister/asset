// Package order allows to calculate future price projections, broker fees and risk figures
package order

import (
	"github.com/MaximilianMeister/asset/broker"
	"github.com/shopspring/decimal"
)

// Order contains data to calculate order figures
type Order struct {
	brokerAlias          string
	volume               int64
	target, actual, stop decimal.Decimal
}

// RiskRewardRatio returns a risk reward ratio value for an order
func (o *Order) RiskRewardRatio() (rrr decimal.Decimal) {
	chance := o.target.Sub(o.actual)
	risk := o.actual.Sub(o.stop)
	rrr = chance.Div(risk).Round(1)

	return
}

// TotalCommission returns the total broker commission fee for an order
func (o *Order) TotalCommission(brokerAlias string) (commission decimal.Decimal, err error) {
	commission = decimal.NewFromFloat(0.0)

	broker, err := broker.FindBroker(brokerAlias)
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

// StopLoss returns nil when it is valid.
func StopLoss(actual, stop float64) error {
	if stop >= actual {
		return &HigherLowerError{stop, actual}
	}

	return nil
}
