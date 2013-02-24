package binarytree

import (
    "testing"
)

// In this test, we are creating an unbalanced binary tree
func TestTree(t *testing.T) {
    e := New("Root")
    e.PushLeft("Root:Left")
    e.PushRight("Root:Right")
    e.Left().PushLeft("Root:Left:Left")
    e.Left().PushRight("Root:Left:Right")
    e.Left().Right().PushRight("Root:Left:Right:Right")

    //Get a channel with our output and print it out line by line
    x := Walker(e)
    val := ""
    for y := range x {
        val += y.(string) + "::"
    }

    expected := "Root::Root:Left::Root:Left:Left::Root:Left:Right::Root:Left:Right:Right::Root:Right::"
    if val != expected {
        t.Errorf("Unexpected value. Received %s, expected %s.", val, expected)
    }
}
