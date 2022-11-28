package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	t.Run("Should be able to observe events published", func(t *testing.T) {
		p := behavioral.NewPatient("Fabrício", 35)
		ds := &behavioral.DoctorService{}
		p.Subscribe(ds)

		p.CatchACold()
		lastMsg := ds.LastMessage()

		assert.Equal(t, "A doctor has been called for Fabrício", lastMsg)
	})

	t.Run("Should be able to observe property changes", func(t *testing.T) {
		p := behavioral.NewClient(50)
		tm := &behavioral.TrafficManagement{Publisher: p.Publisher}
		p.Subscribe(tm)

		for i := 16; i <= 20; i++ {
			p.SetAge(i)
		}

		assert.Len(t, tm.Messages, 1)
		assert.Equal(t, "Congrats, you can drive now!", tm.Messages[0])
	})
}
