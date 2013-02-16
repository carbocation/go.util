package main

import (
    "fmt"
    "runtime"
)

func main() {
    //Run with max possible CPUs
    parallelism := 4
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    var N int = 1e5 //Number of integers that will be in the channel

    c := make(chan int, parallelism) //Create our channel

    //In a separate thread, fill the channel
    go func() {
        defer close(c)
        fillChannel(c, N)
    }()

    //In this thread, drain the channel as it fills
    func() {
        drainChannel(c)
    }()

    fmt.Println(parallelism)
}

func fillChannel(c chan int, N int) {
    for i := 0; i < N; i++ {
        fmt.Println("Filling channel")
        c <- i
    }
}

func drainChannel(c chan int) {
    for i := range c {
        fmt.Println("Draining channel of",i)
    }
}