/*
Create two slices of integers.

Then, figure out which of those integers are in both lists.
*/
package main

import (
	"fmt"
	"math/rand"

//    "sync"
)

var length int = 1000
var domain int = 100000

var A []int = make([]int, length)
var B []int = make([]int, length)

//var C [][]int = make([][]int, 0, length)

var ch chan []int = make(chan []int, 1)

//Generate random values
func construct(Group []int) {
	for i := range Group {
		Group[i] = rand.Intn(domain)
	}
}

func main() {
	construct(A)
	construct(B)

	//var wg sync.WaitGroup

	for ak, av := range A {
		//    wg.Add(1)

		//    go func(wg *sync.WaitGroup) {
		//        defer wg.Done()
		ak, av := ak, av

		for bk, bv := range B {
			if av == bv {
				ch <- []int{ak, bk, av}
			}
		}
		//    }(&wg)
	}

	go func() {
		for i := range ch {
			//C = append(C, []int{ak, bk, av})
			fmt.Println(i)
		}
	}()
	/*
	   go func(wg *sync.WaitGroup) {
	       wg.Wait()
	       defer close(ch)
	   }(&wg)
	*/
	fmt.Println(A)
	fmt.Println(B)
	//fmt.Println(C)
}
