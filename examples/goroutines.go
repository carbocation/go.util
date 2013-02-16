package main

import (
    "fmt"
    "runtime"
)

func main() {
    //Run with max possible CPUs
    numCpu := runtime.NumCPU()
    runtime.GOMAXPROCS(numCpu)
    
    //Number of integers that will be in the channel
    var N int = 1e7

    //Create our channel with enough space to hold all integers at once
    c := make(chan int, N)

    //In a separate thread, fill the channel
    go func() {
        defer close(c)
        start := 0
        stop := N
        fillChannel(c, start, stop)
    }()

    //In this thread, drain the channel
    sum := drainChannel(c)

    fmt.Println("By manual counting, the sum of all integers from 0 to",(N-1),"is",sum)
}

func fillChannel(c chan int, start int, stop int) {
    for i := start; i < stop; i++ {
        c <- i
    }
}

func drainChannel(c chan int) int64 {
    var x int64 = 0
    for i := range c {
        x += int64(i)
    }

    return x
}