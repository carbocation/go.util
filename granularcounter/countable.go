package granularcounter

type Countable interface {
	Count() int
	Add(int)
}

type countable struct {
	name int
	val  int
}

func (c countable) Count() int {
	return c.val
}

func (c *countable) Add(v int) {
	c.val += v
}

func MakeCountable(v int) *countable {
	return &countable{val: v}
}
