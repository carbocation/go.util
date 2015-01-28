package granularcounter

import (
	"sync"

	"github.com/zfjagann/golang-ring"
)

func NewGranularCounter(bufferCap int) *GranularCounter {
	g := &GranularCounter{
		buffer:     &ring.Ring{},
		bufferCap:  bufferCap,
		bufferUsed: 0,
		lastName:   -1,
	}

	g.buffer.SetCapacity(bufferCap)

	return g
}

type GranularCounter struct {
	parent *GranularCounter
	child  *GranularCounter

	buffer     *ring.Ring
	bufferCap  int // Total slots
	bufferUsed int // Non-empty slots

	//The most recent chunk encountered by an Add call
	lastName int

	// Dictates when the child's buffer should be accumulated into a buffer of this counter
	nameFunc func() int

	sync.RWMutex
}

func (g *GranularCounter) Values() []interface{} {
	return g.buffer.Values()
}

func (g *GranularCounter) SumChildren() int {
	g.RLock()
	defer g.RUnlock()

	sum := g.Sum()
	if g.child != nil {
		sum += g.child.SumChildren()
	}

	return sum
}

func (g *GranularCounter) Len() int {
	g.RLock()
	defer g.RUnlock()

	return len(g.buffer.Values())
}

func (g *GranularCounter) Sum() int {
	g.RLock()
	defer g.RUnlock()

	sum := 0
	for _, val := range g.buffer.Values() {
		sum += val.(Countable).Count()
	}

	return sum
}

func (g *GranularCounter) NewParent(nameFunc func() int, bufferCap int) *GranularCounter {
	parent := NewGranularCounter(bufferCap)
	parent.nameFunc = nameFunc
	g.parent = parent
	parent.child = g

	return parent
}

func (g *GranularCounter) Add(v Countable) {
	g.Lock()
	defer g.Unlock()

	if g.parent != nil {
		// If the parent's naming function yields a new result, we can shove
		// the current data into the parent's data buffer and empty our buffer
		if name := g.parent.nameFunc(); name != g.lastName {
			g.Unlock()
			g.parent.Add(&countable{name: name, val: g.Sum()})
			g.Lock()

			g.lastName = name

			// Reset the buffer, but preserve its capacity
			g.buffer = &ring.Ring{}
			g.buffer.SetCapacity(g.bufferCap)
			g.bufferUsed = 1

			// Add the new value
			g.buffer.Enqueue(v)

			return
		} else if g.bufferUsed == g.bufferCap {
			// If it is not time to roll up the data into the parent, but the
			// buffer is at capacity, we will start adding to minimize loss
			g.Unlock()
			elem := g.buffer.Dequeue().(Countable)
			elem.Add(v.Count())
			g.buffer.Enqueue(elem)
			g.Lock()

			return
		}
		// If the buffer is not full and the parent's naming function does not
		// yield a new result, continue with the usual enqueue call
	}

	// We only roll up into a parent buffer if there is a parent
	// Otherwise we continuously add to a ring buffer
	g.buffer.Enqueue(v)
	if g.bufferUsed < g.bufferCap {
		g.bufferUsed++
	}

}
