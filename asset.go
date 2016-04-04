package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Broker struct {
	Name           string  `json:"name"`
	BasicPrice     float64 `json:"basic_price"`
	CommissionRate float64 `json:"commission_rate"`
	MinRate        float64 `json:"min_rate"`
	MaxRate        float64 `json:"max_rate"`
}

type Brokers map[string]Broker

type Order struct {
	brokerAlias          string
	volume               uint32
	target, actual, stop float64
}

func BrokerRegister() (brokers Brokers, err error) {
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

func StopLoss(actual, stop float64) (float64, error) {
	if stop >= actual {
		return actual, &higherLowerError{stop, actual}
	}

	return stop, nil
}

func RiskRewardRatio(o Order) (rrr float64) {
	chance := o.target - o.actual
	risk := o.actual - o.stop
	rrr = RoundDown(float64(chance/risk), 1)

	return
}

func TotalCommission(o Order, brokerAlias string) (commission float64, err error) {
	commission = 0.0

	broker, err := FindBroker(brokerAlias)
	if err != nil {
		return commission, err
	}

	volumeRateBuy := float64(Amount(o)) * o.actual * broker.CommissionRate
	volumeRateSell := float64(Amount(o)) * o.target * broker.CommissionRate

	buySell := []float64{
		volumeRateBuy,
		volumeRateSell,
	}

	for _, bs := range buySell {
		if (bs + broker.BasicPrice) > broker.MinRate {
			if (bs + broker.BasicPrice) > broker.MaxRate {
				commission += broker.MaxRate
			} else {
				commission += (broker.BasicPrice + bs)
			}
		} else {
			commission += broker.MinRate
		}
	}
	commission = RoundUp(float64(commission), 2)

	return
}

func Amount(o Order) (amount uint32) {
	amountFloat := float64(o.volume) / o.actual
	amount = uint32(RoundDown(float64(amountFloat), 0))

	return
}

func Gain(o Order, broker string) (gain float64, err error) {
	amount := Amount(o)
	commission, err := TotalCommission(o, broker)
	if err != nil {
		return 0.0, err
	}

	gain = (float64(amount) * o.target) - float64(o.volume) - commission
	gain = RoundUp(float64(gain), 2)

	return gain, nil
}

func Loss(o Order, broker string) (loss float64, err error) {
	amount := Amount(o)
	commission, err := TotalCommission(o, broker)
	if err != nil {
		return 0.0, err
	}

	loss = float64(o.volume) - (float64(amount) * o.stop) + commission
	loss = RoundUp(float64(loss), 2)

	return loss, nil
}

func Even(o Order, broker string) (even float64, err error) {
	amount := Amount(o)
	commission, err := TotalCommission(o, broker)
	if err != nil {
		return 0.0, err
	}

	even = (float64(o.volume) + commission) / float64(amount)
	even = RoundUp(float64(even), 2)

	return even, nil
}
