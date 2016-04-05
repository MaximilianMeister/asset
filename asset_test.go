package asset

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestStopLoss(t *testing.T) {
	for _, n := range stopLossTests {
		sl, err := StopLoss(n.actual, n.stop)
		if err != nil && n.errExpected == false {
			t.Error(err)
		}
		if sl != n.expected {
			t.Errorf("%f should be %f", sl, n.expected)
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
		t.Errorf("%s is not a Broker register", b)
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
			t.Errorf("%t should be %t", b, n.expected)
		}
	}
	t.Log(len(isBrokerTests), "test cases")
}

func TestFindBroker(t *testing.T) {
	success := true
	for _, n := range findBrokerTests {
		b, err := FindBroker(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if b.Name != n.expected.Name {
			success = false
		}
		if !b.BasicPrice.Equals(n.expected.BasicPrice) {
			success = false
		}
		if !b.CommissionRate.Equals(n.expected.CommissionRate) {
			success = false
		}
		if !b.MinRate.Equals(n.expected.MinRate) {
			success = false
		}
		if !b.MaxRate.Equals(n.expected.MaxRate) {
			success = false
		}
		if !success {
			t.Errorf("%s should be %s", b, n.expected)
		}
	}
	t.Log(len(findBrokerTests), "test cases")
}

func TestCreateOrder(t *testing.T) {
	for _, n := range orderTests {
		o := Order{
			n.brokerAlias,
			n.volume,
			decimal.NewFromFloat(n.target),
			decimal.NewFromFloat(n.actual),
			decimal.NewFromFloat(n.stop),
		}
		if reflect.TypeOf(o).String() != "asset.Order" {
			t.Error("Not of type Order")
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestRiskRewardRatio(t *testing.T) {
	for _, n := range orderTests {
		o := Order{
			n.brokerAlias,
			n.volume,
			decimal.NewFromFloat(n.target),
			decimal.NewFromFloat(n.actual),
			decimal.NewFromFloat(n.stop),
		}
		rrr := o.RiskRewardRatio()
		if !rrr.Equals(n.rrr) {
			t.Errorf("%v should be %v", rrr, n.rrr)
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestTotalCommission(t *testing.T) {
	for _, n := range orderTests {
		o := Order{
			n.brokerAlias,
			n.volume,
			decimal.NewFromFloat(n.target),
			decimal.NewFromFloat(n.actual),
			decimal.NewFromFloat(n.stop),
		}
		totalCommission, err := o.TotalCommission(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if !totalCommission.Equals(decimal.NewFromFloat(n.commission)) {
			t.Errorf("%v should be %v", totalCommission, n.commission)
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestAmount(t *testing.T) {
	for _, n := range orderTests {
		o := Order{
			n.brokerAlias,
			n.volume,
			decimal.NewFromFloat(n.target),
			decimal.NewFromFloat(n.actual),
			decimal.NewFromFloat(n.stop),
		}
		amount := o.Amount()
		if amount != n.amount {
			t.Errorf("%d should be %d", amount, n.amount)
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestGain(t *testing.T) {
	for _, n := range orderTests {
		o := Order{
			n.brokerAlias,
			n.volume,
			decimal.NewFromFloat(n.target),
			decimal.NewFromFloat(n.actual),
			decimal.NewFromFloat(n.stop),
		}
		gain, err := o.Gain(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if !gain.Equals(decimal.NewFromFloat(n.gain)) {
			t.Errorf("%v should be %v", gain, n.gain)
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestLoss(t *testing.T) {
	for _, n := range orderTests {
		o := Order{
			n.brokerAlias,
			n.volume,
			decimal.NewFromFloat(n.target),
			decimal.NewFromFloat(n.actual),
			decimal.NewFromFloat(n.stop),
		}
		loss, err := o.Loss(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if !loss.Equals(decimal.NewFromFloat(n.loss)) {
			t.Errorf("%v should be %v", loss, n.loss)
		}
	}
	t.Log(len(orderTests), "test cases")
}

func TestEven(t *testing.T) {
	for _, n := range orderTests {
		o := Order{
			n.brokerAlias,
			n.volume,
			decimal.NewFromFloat(n.target),
			decimal.NewFromFloat(n.actual),
			decimal.NewFromFloat(n.stop),
		}
		even, err := o.Even(n.brokerAlias)
		if err != nil {
			t.Error(err)
		}
		if !even.Equals(decimal.NewFromFloat(n.even)) {
			t.Errorf("%v should be %v", even, n.even)
		}
	}
	t.Log(len(orderTests), "test cases")
}
