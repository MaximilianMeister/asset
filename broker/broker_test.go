package broker

import (
	"testing"

	"github.com/franela/goblin"
)

func TestNewBrokers(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("New Brokers", func() {
		register, err := NewBrokers()
		for i, n := range newBrokersTests {
			g.It("Should Return A Map Of Brokers", func() {
				g.Assert(err == nil).IsTrue()
				g.Assert(register[i] == n)
				g.Assert(register[i].BasicPrice == n.BasicPrice)
				g.Assert(register[i].CommissionRate == n.CommissionRate)
				g.Assert(register[i].MinRate == n.MinRate)
				g.Assert(register[i].MaxRate == n.MaxRate)
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
