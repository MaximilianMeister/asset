package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Broker struct {
	Name           string  `json:"name"`
	BasicPrice     float32 `json:"basic_price"`
	CommissionRate float32 `json:"commission_rate"`
	MinRate        float32 `json:"min_rate"`
	MaxRate        float32 `json:"max_rate"`
}

type Brokers map[string]Broker

type Order struct {
	broker               string
	volume               uint32
	target, actual, stop float32
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

func IsBroker(shortName string) bool {
	for b := range BrokerRegister() {
		if b == shortName {
			return true
		}
	}
	return false
}

func GetBroker(shortName string) Broker {
	if IsBroker(shortName) {
		register := BrokerRegister()
		for b := range register {
			if b == shortName {
				return register[b]
			}
		}
	}

	return Broker{}
}

func StopLoss(actual, stop float32) float32 {
	if stop >= actual {
		return actual
	}

	return stop
}

func New(o Order) Order {
	return o
}

func RiskRewardRatio(o Order) float32 {
	chance := o.target - o.actual
	risk := o.actual - o.stop
	rrr := chance / risk

	return float32(RoundDown(float64(rrr), 1))
}

func TotalCommission(o Order, shortName string) (commission float32) {
	commission = 0.0

	broker := GetBroker(shortName)
	volumeRateBuy := float32(Amount(o)) * o.actual * broker.CommissionRate
	volumeRateSell := float32(Amount(o)) * o.target * broker.CommissionRate

	buySell := []float32{
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
	commission = float32(RoundUp(float64(commission), 2))

	return
}

func Amount(o Order) uint32 {
	amount := float32(o.volume) / o.actual
	amountRounded := RoundDown(float64(amount), 0)

	return uint32(amountRounded)
}

func Gain(o Order, broker string) float32 {
	amount := Amount(o)
	commission := TotalCommission(o, broker)
	gain := (float32(amount) * o.target) - float32(o.volume) - commission

	return float32(RoundUp(float64(gain), 2))
}

func Loss(o Order, broker string) float32 {
	amount := Amount(o)
	commission := TotalCommission(o, broker)
	loss := float32(o.volume) - (float32(amount) * o.stop) + commission

	return float32(RoundUp(float64(loss), 2))
}

func Even(o Order, broker string) float32 {
	amount := Amount(o)
	commission := TotalCommission(o, broker)
	even := (float32(o.volume) + commission) / float32(amount)

	return float32(RoundUp(float64(even), 2))
}
