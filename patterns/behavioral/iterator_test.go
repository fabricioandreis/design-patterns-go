package behavioral_test

import (
	"testing"

	"github.com/fabricioandreis/design-patterns-go/patterns/behavioral"
	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	t.Run("Should be able to iterate with an array", func(t *testing.T) {
		p := behavioral.Person{"Alexander", "Graham", "Bell"}

		array := p.Names()

		assert.Len(t, array, 3)
		for _, name := range array {
			assert.NotEqual(t, "", name)
		}
	})

	t.Run("Should be able to iterate with a generator", func(t *testing.T) {
		p := behavioral.Person{"Alexander", "Graham", "Bell"}

		generator := p.NamesGenerator()

		for name := range generator {
			assert.NotEqual(t, "", name)
		}
	})

	t.Run("Should be able to iterate with a iterator", func(t *testing.T) {
		// This is the non-idiomatic way (similar to other languages like Java and C++)
		p := behavioral.Person{"Alexander", "Graham", "Bell"}

		for it := behavioral.NewPersonNameIterator(&p); it.MoveNext(); {
			assert.NotEqual(t, "", it.Value())
		}
	})

	buildBinaryTreeRoot := func() *behavioral.Node {
		//   1
		//  / \
		// 2   3
		// Traversal algorithms:
		// - in-order: 2 1 3
		// - pre-order: 1 2 3
		// - post-order: 2 3 1

		root := behavioral.NewNode(
			1,
			behavioral.NewLeafNode(2),
			behavioral.NewLeafNode(3),
		)
		return root
	}

	t.Run("Should be able to iterate with a iterator with complex structures", func(t *testing.T) {
		root := buildBinaryTreeRoot()
		expected := []int{2, 1, 3}

		output := []int{}
		for it := behavioral.NewInOrderIterator(root); it.MoveNext(); {
			output = append(output, it.Value())
		}

		assert.NotNil(t, root)
		assert.Len(t, output, len(expected))
		for i := range expected {
			assert.Equal(t, expected[i], output[i])
		}
	})

	t.Run("Should be able to iterate with a iterator over a binary tree", func(t *testing.T) {
		root := buildBinaryTreeRoot()
		tree := behavioral.NewBinaryTree(root)
		expected := []int{2, 1, 3}

		output := []int{}
		for it := tree.InOrder(); it.MoveNext(); {
			output = append(output, it.Value())
		}

		assert.NotNil(t, root)
		assert.Len(t, output, len(expected))
		for i := range expected {
			assert.Equal(t, expected[i], output[i])
		}
	})
}
