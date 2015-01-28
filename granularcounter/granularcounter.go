package granularcounter

import (
	"sync"
	"time"

	"github.com/zfjagann/golang-ring"
)

func NewGranularCounter(nameFunc func() int, sliverCap int) *granularCounter {
	g := &granularCounter{
		slivers:  &ring.Ring{},
		nameFunc: nameFunc,
		lastName: -1,
	}

	g.slivers.SetCapacity(sliverCap)

	return g
}

type granularCounter struct {
	parent *granularCounter
	child  *granularCounter

	slivers *ring.Ring

	//The most recent chunk encountered by an Add call
	lastName int

	// Dictates when the slivers should be accumulated into a slice of the parent
	nameFunc func() int

	sync.RWMutex
}

func (g *granularCounter) SumChildren() int {
	g.RLock()
	defer g.RUnlock()

	sum := g.Sum()
	if g.child != nil {
		sum += g.child.SumChildren()
	}

	return sum
}

func (g *granularCounter) Len() int {
	g.RLock()
	defer g.RUnlock()

	return len(g.slivers.Values())
}

func (g *granularCounter) Sum() int {
	g.RLock()
	defer g.RUnlock()

	sum := 0
	for _, val := range g.slivers.Values() {
		sum += val.(int)
	}

	return sum
}

func (g *granularCounter) NewParent(nameFunc func() int, sliverCap int) *granularCounter {
	parent := NewGranularCounter(nameFunc, sliverCap)
	g.parent = parent
	parent.child = g

	return parent
}

func (g *granularCounter) Add(v int) {
	g.Lock()
	defer g.Unlock()

	if name := g.nameFunc(); name != g.lastName && g.parent != nil {
		g.lastName = name

		g.Unlock()
		g.parent.Add(g.Sum())
		g.Lock()

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

func nanosecond() int {
	t := time.Now()

	return t.Nanosecond()
}

func microsecond() int {
	t := time.Now()

	return t.Nanosecond() / 1000
}

func millisecond() int {
	t := time.Now()

	return t.Nanosecond() / 1000000
}

func second() int {
	t := time.Now()

	return t.Second()
}

func minute() int {
	t := time.Now()

	return t.Minute()
}
