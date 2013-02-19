/*
Maporder demonstrates that ranging over maps goes in a random order. Specifically, 
and I'm sure this is undefined and subject to change, it selects map elements in 
a uniformly random order.
*/
package main

import (
    "fmt"
)

func main() {
    Ntrials := 10000
    N := 100
    hist := make(map[int]int, N)
    binom := map[bool]int{true: 0, false: 0}

    trials := make([][]int, Ntrials)

    for i := 0; i < Ntrials; i++ {
        trials[i] = trial(N)
    }

    for i := 0; i < N; i++ {
        hist[i] = 0
    }
    
    for _, trial := range trials {
        for j, value := range trial {
            hist[j] += value
        }
    }
    
    for i, v := range hist {
        hist[i] = int(float64(v) / float64(N/2))
        if hist[i] > Ntrials {
            binom[true] ++
        }else{
            binom[false] ++
        }
    }

    fmt.Println(hist)
    fmt.Println(binom)
}

func trial(N int) []int {
    m := make(map[int]string, N)
    l := make([]int, N)

    //Construct the map
    for i := 0; i < N; i++ {
        m[i] = "A"
    }
    
    i := 0
    for id, _ := range m {
        l[i] = id
        i++
    }

    return l
}