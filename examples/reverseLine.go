package main

import (
	"bufio"
	"fmt"
	"github.com/carbocation/util.git/str"
	"os"
)

//Currently reverses ONLY THE FIRST LINE PASSED TO IT
//@TODO@ Allow it to read an indefinite number of lines, reversing each in turn
//@TODO@ Use goroutines or channels or whatever to concurrently read in lines
// and output reversed lines instead of doing these things sequentially, using
// multiple cores
func main() {
	stdIn := bufio.NewReader(os.Stdin)
	input, _ := stdIn.ReadString('\n')

	fmt.Printf("%s", str.Reverse(input))
}
