//From http://rosettacode.org/wiki/Stack#Go 2013/01/22
package datatypes

type Stack []interface{}
 
func (k *Stack) Push(s interface{}) {
    *k = append(*k, s)
}
 
func (k *Stack) Pop() (s interface{}, ok bool) {
    if k.Empty() {
        return
    }
    last := len(*k) - 1
    s = (*k)[last]
    *k = (*k)[:last]
    return s, true
}
 
func (k *Stack) peek() (s interface{}, ok bool) {
    if k.Empty() {
        return
    }
    last := len(*k) - 1
    s = (*k)[last]
    return s, true
}
 
func (k *Stack) Empty() bool {
    return len(*k) == 0
}

/*
func main() {
    s := Stack{}
    s.Push("hi")
    s.Push("there")
    s.Push("bruh")
    
    for !s.Empty() {
        l, ok := s.Pop()
        if ok {
            fmt.Println(l)
        }
    }
}
*/