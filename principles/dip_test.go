package principles_test

import (
	"testing"

	dip "github.com/fabricioandreis/design-patterns-go/principles"
	"github.com/stretchr/testify/assert"
)

func TestDIPViolated(t *testing.T) {
	parent := dip.Person{"Fabrício"}
	child := dip.Person{"Rafael"}
	rels := dip.Relationships{}
	rels.AddParentAndChild(&parent, &child)

	r := dip.Research{rels}
	isParent := r.Investigate("Fabrício", dip.Parent)
	isChild := r.Investigate("Rafael", dip.Child)

	assert.Equal(t, true, isParent)
	assert.Equal(t, true, isChild)
}

func TestDIPAdhered(t *testing.T) {
	parent := dip.Person{"Daiana"}
	child := dip.Person{"Rafael"}
	rels := dip.RelationshipsDIP{}
	rels.AddParentAndChild(&parent, &child)

	r := dip.ResearchDIP{&rels}
	children := r.AllChildren("Daiana")

	assert.Equal(t, 1, len(children))
	assert.NotNil(t, children[0])
	assert.Equal(t, children[0], &child)
	assert.Equal(t, children[0].Name, child.Name)
}
