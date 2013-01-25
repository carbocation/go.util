package str

//Reverses a string
//Inspired by http://golang.org/doc/effective_go.html#for
func Reverse(s string) string{
    //Initialize our output as a byte-ified version of string s
    o := []rune(s)
    
    //Loop over the string, swapping the Nth and the (len(s)-1 - N)th 
    // elements until you reach the midpoint.
    for i, j := 0, len(o)-1; i < j; i, j = i+1, j-1 {
        o[i], o[j] = o[j], o[i]
    }

    return string(o)
}
