//Derived from the stack implementation at http://rosettacode.org/wiki/Queue#Go 2013/01/22
package datatypes

type Queue []interface{}
 
func (k *Queue) Push(s interface{}) {
    *k = append(*k, s)
}
 
func (k *Queue) Pop() (s interface{}, ok bool) {
    if k.Empty() {
        return
    }
    s = (*k)[0]
    *k = (*k)[1:]
    return s, true
}
 
func (k *Queue) peek() (s interface{}, ok bool) {
    if k.Empty() {
        return
    }
    last := len(*k) - 1
    s = (*k)[last]
    return s, true
}
 
func (k *Queue) Empty() bool {
    return len(*k) == 0
}

/*
func main() {
    s := Queue{}
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