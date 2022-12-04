package behavioral_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestChainOfResponsibility(t *testing.T) {
	t.Run("Should create a creature", func(t *testing.T) {
		c := behavioral.NewCreature("monster", 100, 56)

		assert.Equal(t, "monster", c.Name)
		assert.Equal(t, 100, c.Attack)
		assert.Equal(t, 56, c.Defense)
	})

	t.Run("Should call a chain of methods", func(t *testing.T) {
		goblin := behavioral.NewCreature("Goblin", 1, 1)

		root := behavioral.NewCreatureModifier(goblin)
		root.Add(behavioral.NewDoubleAttackModifier(goblin))
		root.Add(behavioral.NewIncreaseDefenseModifier(goblin))
		root.Add(behavioral.NewDoubleAttackModifier(goblin))
		root.Handle()

		assert.Equal(t, "Goblin", goblin.Name)
		assert.Equal(t, 4, goblin.Attack)
		assert.Equal(t, 2, goblin.Defense)
		fmt.Println(goblin)
	})

	t.Run("Should be able to cancel chain execution", func(t *testing.T) {
		goblin := behavioral.NewCreature("Goblin", 1, 1)

		root := behavioral.NewCreatureModifier(goblin)
		root.Add(behavioral.NewNoBonusesModifier(goblin))
		root.Add(behavioral.NewDoubleAttackModifier(goblin))
		root.Add(behavioral.NewIncreaseDefenseModifier(goblin))
		root.Add(behavioral.NewDoubleAttackModifier(goblin))
		root.Handle()

		assert.Equal(t, "Goblin", goblin.Name)
		assert.Equal(t, 1, goblin.Attack)
		assert.Equal(t, 1, goblin.Defense)
		fmt.Println(goblin)
	})

	t.Run("Should be able to implement chain of responsibility, mediator, observer, and CQS", func(t *testing.T) {
		game := &behavioral.BrokerGame{sync.Map{}}

		goblin := behavioral.NewBrokerCreature(game, "Goblin", 2, 2)
		fmt.Println(goblin)
		m := behavioral.NewBrokerDoubleAttachModifier(game, goblin)

		assert.Equal(t, "Goblin", goblin.Name)
		assert.Equal(t, 4, goblin.Attack())
		assert.Equal(t, 2, goblin.Defense())
		fmt.Println(goblin)
		m.Close() // Unsubscribes the modifier from any handlers of the Attack value Query
		assert.Equal(t, 2, goblin.Attack())
		assert.Equal(t, 2, goblin.Defense())
		fmt.Println(goblin)
	})
}
