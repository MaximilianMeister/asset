package asset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStopLoss(t *testing.T) {
	for _, n := range stopLossTests {
		sl, err := StopLoss(n.actual, n.stop)
		if err != nil && n.errExpected == false {
			t.Error(err)
		}
		if sl != n.expected {
			t.Error(fmt.Sprintf("%f should be %f", sl, n.expected))
		}
	}
	t.Log(len(stopLossTests), "test cases")
}

func TestBrokerRegister(t *testing.T) {
	b, err := BrokerRegister()
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(b).String() != "asset.Brokers" {
		t.Error(fmt.Sprintf("%s is not a Broker register", b))
	}
	t.Log("1 test cases")
}

func TestIsBroker(t *testing.T) {
	for _, n := range isBrokerTests {
		b, err := IsBroker(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if b != n.expected {
			t.Error(fmt.Sprintf("%t should be %t", b, n.expected))
		}
	}
	t.Log(len(isBrokerTests), "test cases")
}

func TestFindBroker(t *testing.T) {
	for _, n := range findBrokerTests {
		b, err := FindBroker(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if b != n.expected {
			t.Error(fmt.Sprintf("%s should be %s", reflect.TypeOf(b).String(), n.expected))
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
		rrr := o.RiskRewardRatio()
		if rrr != n.rrr {
			t.Error(fmt.Sprintf("%f should be %f", rrr, n.rrr))
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestTotalCommission(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		totalCommission, err := o.TotalCommission(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if totalCommission != n.commission {
			t.Error(fmt.Sprintf("%f should be %f", totalCommission, n.commission))
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestAmount(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		amount := o.Amount()
		if amount != n.amount {
			t.Error(fmt.Sprintf("%d should be %d", amount, n.amount))
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestGain(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		gain, err := o.Gain(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if gain != n.gain {
			t.Error(fmt.Sprintf("%f should be %f", gain, n.gain))
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestLoss(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		loss, err := o.Loss(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if loss != n.loss {
			t.Error(fmt.Sprintf("%f should be %f", loss, n.loss))
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestEven(t *testing.T) {
	for _, n := range orderTests {
		o := Order{n.brokerAlias, n.volume, n.target, n.actual, n.stop}
		even, err := o.Even(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if even != n.even {
			t.Error(fmt.Sprintf("%f should be %f", even, n.even))
		}
	}
	t.Log(len(orderTests), "test cases")
}
