package asset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStopLoss(t *testing.T) {
	for _, n := range stopLossTests {
		sl, err := StopLoss(n.actual, n.stop);
		if err != nil && n.errExpected == false {
			t.Error(err)
		}
		if sl != n.expected {
			t.Error(fmt.Sprintf("%s should be %s", sl, n.expected))
		}
	}
	t.Log(len(stopLossTests), "test cases")
}

func TestBrokerRegister(t *testing.T) {
	b := BrokerRegister()
	if reflect.TypeOf(b).String() != "asset.Brokers" {
		t.Error("Can't open Broker register")
	}
	t.Log("1 test cases")
}

func TestIsBroker(t *testing.T) {
	for _, n := range isBrokerTests {
		b := IsBroker(n.brokerAlias)
		if b != n.expected {
			t.Error("Can't validate Broker")
		}
	}
	t.Log(len(isBrokerTests), "test cases")
}

func TestFindBroker(t *testing.T) {
	for _, n := range findBrokerTests {
		b := FindBroker(n.brokerAlias)
		if b != n.expected {
			t.Error("Can't get Broker")
		}
	}
	t.Log(len(findBrokerTests), "test cases")
}

func TestCreateOrder(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		if reflect.TypeOf(o).String() != "asset.Order" {
			t.Error("Not of type Order")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestRiskRewardRatio(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		rrr := RiskRewardRatio(o)
		if reflect.TypeOf(rrr).String() != "float64" {
			t.Error("Not of type float64")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestTotalCommission(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		totalCommission := TotalCommission(o, n.brokerAlias)
		if totalCommission != n.commission {
			t.Error("Failed to determine total commision")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestAmount(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		if Amount(o) != n.amount {
			t.Error("Failed to determine total amount")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestGain(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		if Gain(o, n.brokerAlias) != n.gain {
			t.Error("Failed to determine maximum gain")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestLoss(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		if Loss(o, n.brokerAlias) != n.loss {
			t.Error("Failed to determine maximum loss")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestEven(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		if Even(o, n.brokerAlias) != n.even {
			t.Error("Failed to determine break even")
		}
	}
	t.Log(len(orderTests), "test cases")
}
