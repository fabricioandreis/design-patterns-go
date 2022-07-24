package principles_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/principles"
	"github.com/stretchr/testify/assert"
)

func TestOldFashionedPrinter(t *testing.T) {
	o := principles.OldFashionedPrinter{}

	assert.Panics(t, func() { o.Scan(principles.Document{}) })
}
