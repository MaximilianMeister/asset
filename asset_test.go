package asset

import (
	"reflect"
	"testing"

	"github.com/franela/goblin"
	"github.com/shopspring/decimal"
)

func TestBrokerRegister(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Broker Register", func() {
		b, err := BrokerRegister()
		for i, n := range brokerRegisterTests {
			g.It("Should Return A Map Of Brokers", func() {
				g.Assert(err == nil).IsTrue()
				g.Assert(b[i] == n)
				g.Assert(b[i].BasicPrice == n.BasicPrice)
				g.Assert(b[i].CommissionRate == n.CommissionRate)
				g.Assert(b[i].MinRate == n.MinRate)
				g.Assert(b[i].MaxRate == n.MaxRate)
			})
		}
	})
}

func TestIsBroker(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Is Broker", func() {
		for _, n := range isBrokerTests {
			g.It("Should Be A Broker", func() {
				err := IsBroker(n.brokerAlias)
				if n.errExpected == true {
					g.Assert(err == nil).IsFalse()
				} else {
					g.Assert(err == nil).IsTrue()
				}
			})
		}
	})
}

func TestFindBroker(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Find Broker", func() {
		g.It("Should Find A Broker", func() {
			for _, n := range findBrokerTests {
				b, err := FindBroker(n.brokerAlias)
				g.Assert(err == nil).IsTrue()
				g.Assert(b.Name == n.expected.Name)
				g.Assert(b.BasicPrice == n.expected.BasicPrice)
				g.Assert(b.CommissionRate == n.expected.CommissionRate)
				g.Assert(b.MinRate == n.expected.MinRate)
				g.Assert(b.MaxRate == n.expected.MaxRate)
			}
		})
	})
}

func TestCreateOrder(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Create Order", func() {
		g.It("Should Create An Order", func() {
			for _, n := range orderTests {
				o := Order{
					n.brokerAlias,
					n.volume,
					decimal.NewFromFloat(n.target),
					decimal.NewFromFloat(n.actual),
					decimal.NewFromFloat(n.stop),
				}
				g.Assert(reflect.TypeOf(o).String()).Equal("asset.Order")
			}
		})
	})
}

func TestRiskRewardRatio(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Risk Reward Ratio", func() {
		g.It("Should Return A Risk Reward Ratio", func() {
			for _, n := range orderTests {
				o := Order{
					n.brokerAlias,
					n.volume,
					decimal.NewFromFloat(n.target),
					decimal.NewFromFloat(n.actual),
					decimal.NewFromFloat(n.stop),
				}
				rrr := o.RiskRewardRatio()
				g.Assert(rrr).Equal(n.rrr)
			}
		})
	})
}

func TestTotalCommission(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Total Commission", func() {
		g.It("Should Return A Total Commission", func() {
			for _, n := range orderTests {
				o := Order{
					n.brokerAlias,
					n.volume,
					decimal.NewFromFloat(n.target),
					decimal.NewFromFloat(n.actual),
					decimal.NewFromFloat(n.stop),
				}
				totalCommission, err := o.TotalCommission(n.brokerAlias)
				g.Assert(err == nil).IsTrue()
				g.Assert(totalCommission).Equal(decimal.NewFromFloat(n.commission).Round(2))
			}
		})
	})
}

func TestAmount(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Amount", func() {
		for _, n := range orderTests {
			g.It("Should Return A Total Amount Of Order Assets", func() {
				o := Order{
					n.brokerAlias,
					n.volume,
					decimal.NewFromFloat(n.target),
					decimal.NewFromFloat(n.actual),
					decimal.NewFromFloat(n.stop),
				}
				amount := o.Amount()
				g.Assert(amount).Equal(n.amount)
			})
		}
	})
}

func TestGain(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Gain", func() {
		for _, n := range orderTests {
			g.It("Should Return A Maximum Gain Of An Order", func() {
				o := Order{
					n.brokerAlias,
					n.volume,
					decimal.NewFromFloat(n.target),
					decimal.NewFromFloat(n.actual),
					decimal.NewFromFloat(n.stop),
				}
				gain, err := o.Gain(n.brokerAlias)
				g.Assert(err == nil).IsTrue()
				g.Assert(gain).Equal(decimal.NewFromFloat(n.gain))
			})
		}
	})
}

func TestLoss(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Loss", func() {
		for _, n := range orderTests {
			g.It("Should Return A Maximum Loss Of An Order", func() {
				o := Order{
					n.brokerAlias,
					n.volume,
					decimal.NewFromFloat(n.target),
					decimal.NewFromFloat(n.actual),
					decimal.NewFromFloat(n.stop),
				}
				loss, err := o.Loss(n.brokerAlias)
				g.Assert(err == nil).IsTrue()
				g.Assert(loss).Equal(decimal.NewFromFloat(n.loss))
			})
		}
	})
}

func TestEven(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Even", func() {
		for _, n := range orderTests {
			g.It("Should Return A Break Even Of An Order", func() {
				o := Order{
					n.brokerAlias,
					n.volume,
					decimal.NewFromFloat(n.target),
					decimal.NewFromFloat(n.actual),
					decimal.NewFromFloat(n.stop),
				}
				even, err := o.Even(n.brokerAlias)
				g.Assert(err == nil).IsTrue()
				g.Assert(even).Equal(decimal.NewFromFloat(n.even))
			})
		}
	})
}

func TestStopLoss(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Stop Loss", func() {
		for _, n := range stopLossTests {
			g.It("Should Evaluate A Stop Loss", func() {
				err := StopLoss(n.actual, n.stop)
				if n.errExpected == true {
					g.Assert(err == nil).IsFalse()
				} else {
					g.Assert(err == nil).IsTrue()
				}
			})
		}
	})
}
