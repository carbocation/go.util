package granularcounter

import (
	"sync"

	"github.com/zfjagann/golang-ring"
)

func NewGranularCounter(nameFunc func() int, sliverCap int) *GranularCounter {
	g := &GranularCounter{
		slivers:  &ring.Ring{},
		nameFunc: nameFunc,
		lastName: -1,
	}

	g.slivers.SetCapacity(sliverCap)

	return g
}

type GranularCounter struct {
	parent *GranularCounter
	child  *GranularCounter

	slivers *ring.Ring

	//The most recent chunk encountered by an Add call
	lastName int

	// Dictates when the slivers should be accumulated into a slice of the parent
	nameFunc func() int

	sync.RWMutex
}

func (g *GranularCounter) Values() []interface{} {
	return g.slivers.Values()
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

	return len(g.slivers.Values())
}

func (g *GranularCounter) Sum() int {
	g.RLock()
	defer g.RUnlock()

	sum := 0
	for _, val := range g.slivers.Values() {
		sum += val.(Countable).Count()
	}

	return sum
}

func (g *GranularCounter) NewParent(nameFunc func() int, sliverCap int) *GranularCounter {
	parent := NewGranularCounter(nameFunc, sliverCap)
	g.parent = parent
	parent.child = g

	return parent
}

func (g *GranularCounter) Add(v Countable) {
	g.Lock()
	defer g.Unlock()

	if name := g.nameFunc(); name != g.lastName && g.parent != nil {
		g.Unlock()
		g.parent.Add(countable{name: name, val: g.Sum()})
		g.Lock()

		g.lastName = name

		// Reset the sliver buffer, but preserve its capacity
		cap := g.slivers.Capacity()
		g.slivers = &ring.Ring{}
		g.slivers.SetCapacity(cap)
	}

	// What to do when we go above capacity?
	/*
		if g.slivers.Peek() != nil {
			g.Unlock()

			oldV := g.slivers.Dequeue().(int)
			v += oldV

			g.Lock()
		}
	*/

	g.slivers.Enqueue(v)
}
