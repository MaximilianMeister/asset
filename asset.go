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

func BrokerRegister() (brokers Brokers) {
	file, e := ioutil.ReadFile("./broker.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	json.Unmarshal(file, &brokers)

	return
}

func IsBroker(brokerAlias string) bool {
	for b := range BrokerRegister() {
		if b == brokerAlias {
			return true
		}
	}
	return false
}

func FindBroker(brokerAlias string) Broker {
	if IsBroker(brokerAlias) {
		register := BrokerRegister()
		for b := range register {
			if b == brokerAlias {
				return register[b]
			}
		}
	}

	return Broker{}
}

func StopLoss(actual, stop float64) float64 {
	if stop >= actual {
		return actual
	}

	return stop
}

func New(o Order) Order {
	return o
}

func RiskRewardRatio(o Order) float64 {
	chance := o.target - o.actual
	risk := o.actual - o.stop
	rrr := chance / risk

	return RoundDown(float64(rrr), 1)
}

func TotalCommission(o Order, brokerAlias string) (commission float64) {
	commission = 0.0

	broker := FindBroker(brokerAlias)
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

func Amount(o Order) uint32 {
	amount := float64(o.volume) / o.actual
	amountRounded := RoundDown(float64(amount), 0)

	return uint32(amountRounded)
}

func Gain(o Order, broker string) float64 {
	amount := Amount(o)
	commission := TotalCommission(o, broker)
	gain := (float64(amount) * o.target) - float64(o.volume) - commission

	return RoundUp(float64(gain), 2)
}

func Loss(o Order, broker string) float64 {
	amount := Amount(o)
	commission := TotalCommission(o, broker)
	loss := float64(o.volume) - (float64(amount) * o.stop) + commission

	return RoundUp(float64(loss), 2)
}

func Even(o Order, broker string) float64 {
	amount := Amount(o)
	commission := TotalCommission(o, broker)
	even := (float64(o.volume) + commission) / float64(amount)

	return RoundUp(float64(even), 2)
}
