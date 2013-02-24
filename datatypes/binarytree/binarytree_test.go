package binarytree

import (
    "testing"
    "fmt"
)

func TestTree(t *testing.T) {
    T := New("Root")
    //E := L.PushFront(1)
    //E := L.Root()
    e := T.Root()
    e.PushLeft("Depth1 Left")
    e.PushRight("Depth1 Right")
    e2 := e.Left()
    //e3 := e.Right()
    e2.PushLeft("Depth2/Left Left")
    e2.PushRight("Depth2/Left Right")
    
    fmt.Printf("%v, %v, %v\n",e, e2, T)

    x := Walker(e)
    fmt.Println("TEST")
    for y := range x {
        fmt.Println(y)
    }
    //Walk(e)
    
    /*
    cur := L.Root().Value
    if cur != 0 {
        t.Errorf("L.Next() should yield 0 here; yields %v", cur)
    }
    */
}
