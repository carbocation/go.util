/*
Create two slices of integers.

Then, figure out which of those integers are in both lists.
*/
package main

import (
    "fmt"
    "math/rand"
)

var length int = 1000
var domain int = 100000

var A []int = make([]int, length)
var B []int = make([]int, length)
var C [][]int = make([][]int, 0, length)


//Generate random values
func construct(Group []int) {
	for i := range Group {
        Group[i] = rand.Intn(domain)
    }
}

func main() {
    construct(A)
    construct(B)
    
    for ak, av := range A {
        for bk, bv := range B {
            if av == bv {
                C = append(C, []int{ak, bk, av})
            }
        }
    }

    fmt.Println(A)
    fmt.Println(B)
    fmt.Println(C)
}