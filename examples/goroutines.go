package main

import (
    "fmt"
    //"time"
    "runtime"
)

func main() {
    //Run with 4 CPUs
    parallelism := runtime.GOMAXPROCS(runtime.NumCPU())
    
    //var N int = 3
    c := make(chan int, parallelism)

    go func() {
        defer close(c)

        fmt.Println("Filling channel")
        c <- 3
        fmt.Println("Filling channel")
        c <- 2
        fmt.Println("Filling channel")
        c <- 1
    }()

    func() {
        lcm1 := len(c)
        k := 0
        
        for i := range c {
            k++
            fmt.Println("Draining channel of",i)

            if k == lcm1 {
                //close(c)
            }
        }
    }()

    //time.Sleep(5 * 1e9)
    //defer close(c)
    //fmt.Println(N, c)
}