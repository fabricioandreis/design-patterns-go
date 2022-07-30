package creational_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/creational"
	"github.com/stretchr/testify/assert"
)

func TestPrototypePattern(t *testing.T) {
	t.Run("Should not deep copy (copy objects pointed at) when directly assigning", func(t *testing.T) {
		john := creational.PersonProto{"John", &creational.AddressProto{"123 London Rd", "London", "UK"}, []string{}}

		jane := john
		jane.Name = "Jane"
		jane.Address.StreetAddress = "321 Baker St"

		assert.Equal(t, "Jane", jane.Name)
		assert.Equal(t, "321 Baker St", jane.Address.StreetAddress)
		assert.Equal(t, john.Address, jane.Address)
		assert.Equal(t, john.Address.StreetAddress, jane.Address.StreetAddress)
		assert.Equal(t, john.Address.City, jane.Address.City)
		assert.Equal(t, john.Address.Country, jane.Address.Country)
	})

	t.Run("Should work when manually coping objects pointed at", func(t *testing.T) {
		john := creational.PersonProto{"John", &creational.AddressProto{"123 London Rd", "London", "UK"}, []string{}}

		jane := john
		jane.Name = "Jane"
		jane.Address = &creational.AddressProto{
			StreetAddress: john.Address.StreetAddress,
			City:          john.Address.City,
			Country:       john.Address.Country,
		}
		jane.Address.StreetAddress = "321 Baker St"

		assert.Equal(t, "Jane", jane.Name)
		assert.NotEqual(t, john.Address, jane.Address)
		assert.Equal(t, "321 Baker St", jane.Address.StreetAddress)
		assert.Equal(t, "London", jane.Address.City)
		assert.Equal(t, "UK", jane.Address.Country)
		assert.Equal(t, "123 London Rd", john.Address.StreetAddress)
		assert.Equal(t, "London", john.Address.City)
		assert.Equal(t, "UK", john.Address.Country)
	})

	t.Run("Should perform deep copy with custom methods", func(t *testing.T) {
		john := creational.PersonProto{
			"John",
			&creational.AddressProto{"123 London Rd", "London", "UK"},
			[]string{"Chris", "Matt"}}

		jane := john.DeepCopy()
		jane.Name = "Jane"
		jane.Address.StreetAddress = "321 Baker St"
		jane.Friends = append(jane.Friends, "Angela")

		assert.Equal(t, "Jane", jane.Name)
		assert.NotEqual(t, john.Address, jane.Address)
		assert.Equal(t, "321 Baker St", jane.Address.StreetAddress)
		assert.Equal(t, "London", jane.Address.City)
		assert.Equal(t, "UK", jane.Address.Country)
		assert.Equal(t, "123 London Rd", john.Address.StreetAddress)
		assert.Equal(t, "London", john.Address.City)
		assert.Equal(t, "UK", john.Address.Country)
		assert.Equal(t, 2, len(john.Friends))
		assert.Equal(t, 3, len(jane.Friends))
	})

	t.Run("Should perform clone through serialization", func(t *testing.T) {
		john := creational.PersonProto{
			"John",
			&creational.AddressProto{"123 London Rd", "London", "UK"},
			[]string{"Chris", "Matt"}}

		jane := john.Clone()
		jane.Name = "Jane"
		jane.Address.StreetAddress = "321 Baker St"
		jane.Friends = append(jane.Friends, "Angela")

		assert.Equal(t, "Jane", jane.Name)
		assert.NotEqual(t, john.Address, jane.Address)
		assert.Equal(t, "321 Baker St", jane.Address.StreetAddress)
		assert.Equal(t, "London", jane.Address.City)
		assert.Equal(t, "UK", jane.Address.Country)
		assert.Equal(t, "123 London Rd", john.Address.StreetAddress)
		assert.Equal(t, "London", john.Address.City)
		assert.Equal(t, "UK", john.Address.Country)
		assert.Equal(t, 2, len(john.Friends))
		assert.Equal(t, 3, len(jane.Friends))
	})

	t.Run("Should use prototype factory methods correctly", func(t *testing.T) {
		main := creational.NewMainOfficeEmployee("John", 100)
		aux := creational.NewAuxOfficeEmployee("Jane", 400)

		assert.Equal(t, "John", main.Name)
		assert.Equal(t, "123 East Dr", main.Office.StreetAddress)
		assert.Equal(t, "London", main.Office.City)
		assert.Equal(t, 100, main.Office.Suite)

		assert.Equal(t, "Jane", aux.Name)
		assert.Equal(t, "66 West Dr", aux.Office.StreetAddress)
		assert.Equal(t, "London", aux.Office.City)
		assert.Equal(t, 400, aux.Office.Suite)
	})
}
