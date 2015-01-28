package granularcounter

type Countable interface {
	Count() int
}

type countable struct {
	name int
	val  int
}

func (c countable) Count() int {
	return c.val
}

func MakeCountable(v int) countable {
	return countable{val: v}
}
