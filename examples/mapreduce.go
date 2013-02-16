package main

import (
    "fmt"
    "runtime"
    //"math"
)

func main() {
    //Run with max possible CPUs
    numCpu := runtime.NumCPU()
    runtime.GOMAXPROCS(numCpu)
    
    //Number of integers that will be in the channel
    //We will start at 0, so if N==1 then you will sum 0, if N==3 you will sum 1+2, etc.
    var N int64 = 1e8

    //Create our channel with enough space to hold all integers at once
    c := make(chan int64, N)

    //In a separate thread, fill the channel
    //Instead of deferring a close in the filling channel, we'll explicitly keep 
    //track of the number of expected channel elements so that we can drain them 
    //via for loop rather than range loop.

    //This now lets us create multiple goroutines to fill the channel, instead of 
    //relying on deferred closure from the filling channel which would limit us to 1
    //filling goroutine at a time.

    //This is basically the map step (one mapper)
    var start, stop int64 = 0, 0
    numGoRoutines := int64(numCpu * 4)
    for i := 0; int64(i) < numGoRoutines; i++ {
        start = int64(i) * N/numGoRoutines
        stop = int64(i+1) * N/numGoRoutines

        go func(start int64, stop int64) {
            fmt.Println("Filling channel from",start,"to",stop)
            fillChannel(c, start, stop)
        }(start, stop)
    }
    if stop < N {
        fmt.Println("Finally filling channel from",stop,"to",N)
        go fillChannel(c, stop, N)
    }

    //In this thread, drain the channel (for example, you could print the records here)
    //This is basically the reduce step (one reducer)
    sum := drainChannel(c, N)

    fmt.Println("By manual counting, the sum of all integers from 0 to",(N-1),"is",sum)
}

func fillChannel(c chan int64, start int64, stop int64) {
    for i := start; int64(i) < stop; i++ {
        c <- int64(i)
    }
}

func drainChannel(c chan int64, N int64) int64 {
    var x int64 = 0
    for i := 0; int64(i) < N; i++ {
        this := <-c
        x += int64(this)
        
        /*
        if math.Mod(float64(i),1e3) == 0.0 {
            fmt.Println("Draining",this)
        }
        */
    }

    return x
}