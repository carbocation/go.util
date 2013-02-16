package main

import (
    "fmt"
    "runtime"
)

func main() {
    //Run with max possible CPUs
    parallelism := 4
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    var N int = 1e3 //Number of integers that will be in the channel

    c := make(chan int, parallelism) //Create our channel

    //In a separate thread, fill the channel
    go func() {
        defer close(c)
        fillChannel(c, N)
    }()

    //In this thread, drain the channel
    sum := drainChannel(c)

    fmt.Println("By manual counting, the sum of all integers from 0 to",(N-1),"is",sum)
}

func fillChannel(c chan int, N int) {
    for i := 0; i < N; i++ {
        c <- i
    }
}

func drainChannel(c chan int) int {
    x := 0
    for i := range c {
        x += i
    }

    return x
}