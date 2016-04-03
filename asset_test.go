package asset

import (
	"reflect"
	"testing"
)

func TestStopLoss(t *testing.T) {
	for _, n := range stopLossTests {
		sl := StopLoss(n.actual, n.stop)
		if sl != n.expected {
			t.Error("Can't calculate stop loss")
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
		b := IsBroker(n.shortName)
		if b != n.expected {
			t.Error("Can't validate Broker")
		}
	}
	t.Log(len(isBrokerTests), "test cases")
}

func TestGetBroker(t *testing.T) {
	for _, n := range getBrokerTests {
		b := GetBroker(n.shortName)
		if b != n.expected {
			t.Error("Can't get Broker")
		}
	}
	t.Log(len(getBrokerTests), "test cases")
}

func TestCreateOrder(t *testing.T) {
	for _, n := range orderTests {
		o := New(Order{n.broker, n.volume, n.target, n.actual, n.stop})
		if reflect.TypeOf(o).String() != "asset.Order" {
			t.Error("Not of type Order")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestRiskRewardRatio(t *testing.T) {
	for _, n := range orderTests {
		o := New(Order{n.broker, n.volume, n.target, n.actual, n.stop})
		rrr := RiskRewardRatio(o)
		if reflect.TypeOf(rrr).String() != "float32" {
			t.Error("Not of type float32")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestTotalCommission(t *testing.T) {
	for _, n := range orderTests {
		o := New(Order{n.broker, n.volume, n.target, n.actual, n.stop})
		totalCommission := TotalCommission(o, n.broker)
		if totalCommission != n.commission {
			t.Error("Failed to determine total commision")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestAmount(t *testing.T) {
	for _, n := range orderTests {
		o := New(Order{n.broker, n.volume, n.target, n.actual, n.stop})
		if Amount(o) != n.amount {
			t.Error("Failed to determine total amount")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestGain(t *testing.T) {
	for _, n := range orderTests {
		o := New(Order{n.broker, n.volume, n.target, n.actual, n.stop})
		if Gain(o, n.broker) != n.gain {
			t.Error("Failed to determine maximum gain")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestLoss(t *testing.T) {
	for _, n := range orderTests {
		o := New(Order{n.broker, n.volume, n.target, n.actual, n.stop})
		if Loss(o, n.broker) != n.loss {
			t.Error("Failed to determine maximum loss")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestEven(t *testing.T) {
	for _, n := range orderTests {
		o := New(Order{n.broker, n.volume, n.target, n.actual, n.stop})
		if Even(o, n.broker) != n.even {
			t.Error("Failed to determine break even")
		}
	}
	t.Log(len(orderTests), "test cases")
}
