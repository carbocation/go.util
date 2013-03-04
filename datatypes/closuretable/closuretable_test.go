package closuretable

import (
	"fmt"
	"github.com/carbocation/util.git/datatypes/binarytree"
    //"github.com/carbocation/util.git/datatypes/closuretable"
	"math/rand"
	"testing"
)

func TestClosureConversion(t *testing.T) {
	// Make some sample entries based on a skeleton
	entries := map[int64]int{
		0: 0, 10: 10, 20: 20, 30: 30, 40: 40, 50: 50, 60: 60,
	}

	// Create a closure table to represent the relationships among the entries
	// In reality, you'd probably directly import the closure table data into the ClosureTable class
	closuretable := ClosureTable{Relationship{Ancestor: 0, Descendant: 0, Depth: 0}}
	closuretable.AddChild(Child{Parent: 0, Child: 10})
	closuretable.AddChild(Child{Parent: 0, Child: 20})
	closuretable.AddChild(Child{Parent: 20, Child: 30})
	closuretable.AddChild(Child{Parent: 30, Child: 40})
	closuretable.AddChild(Child{Parent: 20, Child: 50})
    closuretable.AddChild(Child{Parent: 50, Child: 60})


    // Obligatory boxing step
    // Convert to interface type so the generic TableToTree method can be called on these entries
    interfaceEntries := map[int64]interface{}{}
    for k, v := range entries {
        interfaceEntries[k] = v
    }

	//Build a tree out of the entries based on the closure table's instructions.
	tree := walkBody(closuretable.TableToTree(interfaceEntries))
	expected := 210

	if tree != expected {
		t.Errorf("walkBody(tree) yielded %s, expected %s. Have you made a change that caused the iteration order to become indeterminate, e.g., using a map instead of a slice?", tree, expected)
	}
}

func walkBody(el *binarytree.Tree) int {
	if el == nil {
		return 0
	}

	out := 0
	out += el.Value.(int)
	out += walkBody(el.Left())
	out += walkBody(el.Right())

	return out
}

func buildClosureTable(N int) ClosureTable {
	// Create the closure table with a single progenitor
	ct := ClosureTable{Relationship{Ancestor: 0, Descendant: 0, Depth: 0}}

	for i := 1; i < N; i++ {
		// Create a place for entry #i, making it the child of a random entry j<i
		err := ct.AddChild(Child{Parent: rand.Int63n(int64(i)), Child: int64(i)})
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	return ct
}
