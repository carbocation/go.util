package main

import (
    "os"
    "bufio"
    "fmt"
    "github.com/carbocation/util.git/str"
)

func main() {
    stdIn := bufio.NewReader(os.Stdin)
    input, _ := stdIn.ReadString('\n')

    fmt.Printf("%s", str.Reverse(input))
}