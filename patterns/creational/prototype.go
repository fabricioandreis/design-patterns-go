package creational

import (
	"bytes"
	"encoding/gob"
)

// Prototype is a creational design pattern that lets you copy (clone) existing objects without making your code dependent on their classes.
// Requires *deep copy* support
// Can be more convenient by the use of a factory prototype, which is a factory that serves prototypes
// https://refactoring.guru/design-patterns/prototype

type AddressProto struct {
	StreetAddress, City, Country string
}

func (a *AddressProto) DeepCopy() *AddressProto {
	return &AddressProto{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type PersonProto struct {
	Name    string
	Address *AddressProto
	Friends []string
}

func (p *PersonProto) DeepCopy() *PersonProto {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

// Clone copies the whole object through serialization and deserialization
func (p *PersonProto) Clone() *PersonProto {
	b := bytes.Buffer{}
	result := PersonProto{} // new object (the copied content)

	e := gob.NewEncoder(&b)
	e.Encode(p) // encode (serializes) the object into the buffer

	d := gob.NewDecoder(&b)
	d.Decode(&result) // decode (deserializes) the object from the buffer

	return &result
}

// Prototype Factory
type AddressProtoFact struct {
	Suite               int
	StreetAddress, City string
}

type EmployeeProtoFact struct {
	Name   string
	Office AddressProtoFact
}

func (e *EmployeeProtoFact) Clone() *EmployeeProtoFact {
	b := bytes.Buffer{}
	result := EmployeeProtoFact{} // new object (the copied content)

	enc := gob.NewEncoder(&b)
	enc.Encode(e) // encode (serializes) the object into the buffer

	dec := gob.NewDecoder(&b)
	dec.Decode(&result) // decode (deserializes) the object from the buffer

	return &result
}

var mainOffice = EmployeeProtoFact{
	"", AddressProtoFact{0, "123 East Dr", "London"},
}
var auxOffice = EmployeeProtoFact{
	"", AddressProtoFact{0, "66 West Dr", "London"},
}

func newEmployeeProtoFact(proto *EmployeeProtoFact, name string, suite int) *EmployeeProtoFact {
	result := proto.Clone()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *EmployeeProtoFact {
	return newEmployeeProtoFact(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *EmployeeProtoFact {
	return newEmployeeProtoFact(&auxOffice, name, suite)
}
