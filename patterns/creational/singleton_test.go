package creational_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/creational"
	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	t.Run("Should read data from file", func(t *testing.T) {
		data, err := creational.ExportReadData("capitals.txt")

		assert.Nil(t, err)
		assert.Greater(t, len(data), 0)
	})

	t.Run("Should load singleton data", func(t *testing.T) {
		db := creational.GetSingletonDatabase()

		pop := db.GetPopulation("Seoul")

		assert.Equal(t, 17500000, pop)
	})

	t.Run("Should get total population of cities", func(t *testing.T) {
		totalPop := creational.GetTotalPopulation([]string{"Seoul", "Mexico City", "Osaka"})

		assert.Equal(t, 17500000+17400000+16425000, totalPop) // this assertion is dependent on a concrete database
	})

	t.Run("Should inject the database to get population", func(t *testing.T) {
		totalPop := creational.GetTotalPopulationDIP(&creational.FakeDatabase{}, []string{creational.ExportCityNames[0], creational.ExportCityNames[1]})

		assert.Equal(t, 1+2, totalPop)
	})
}
