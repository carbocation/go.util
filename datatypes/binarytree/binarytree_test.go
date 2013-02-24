package binarytree

import (
    "testing"
    "fmt"
)

// In this test, we are creating an unbalanced binary tree
func TestTree(t *testing.T) {
    e := New("Root")
    e.PushLeft("Root:Left")
    e.PushRight("Root:Right")
    e.Left().PushLeft("Root:Left:Left")
    e.Left().PushRight("Root:Left:Right")
    e.Left().Right().PushRight("Root:Left:Right:Right")
    
    fmt.Printf("%v, %v, %v\n",e)

    //Get a channel with our output and print it out line by line
    x := Walker(e)
    val := ""
    for y := range x {
        fmt.Println(y)
        val += y.(string) + "::"
    }

    fmt.Println(val)
}
