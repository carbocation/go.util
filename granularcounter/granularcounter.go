package granularcounter

import (
	"sync"
	"time"

	"github.com/zfjagann/golang-ring"
)

type granularCounter struct {
	chunks  *ring.Ring
	slivers *ring.Ring

	//The most recent chunk encountered by an Add call
	lastChunkName int

	//The number of Adds() accumulated across all chunks
	chunkSum int

	// Dictates when the slivers should be accumulated into the next chunk
	chunkFunc func() int

	sync.RWMutex
}

func NewGranularCounter(chunkFunc func() int, chunkCap, sliverCap int) *granularCounter {
	g := &granularCounter{
		chunks:    &ring.Ring{},
		slivers:   &ring.Ring{},
		chunkFunc: chunkFunc,
	}
	g.lastChunkName = -1

	g.chunks.SetCapacity(chunkCap)
	g.slivers.SetCapacity(sliverCap)

	return g
}

func minute() int {
	t := time.Now()

	return t.Minute()
}

func second() int {
	t := time.Now()

	return t.Second()
}

func nanosecond() int {
	t := time.Now()

	return t.Nanosecond()
}

func (g *granularCounter) Add() {
	g.Lock()
	defer g.Unlock()

	if chunkName := g.chunkFunc(); chunkName != g.lastChunkName {
		chunk := &granule{Name: chunkName, Count: len(g.slivers.Values())}

		g.chunkSum += chunk.Count

		nextSlot, ok := g.chunks.Peek().(*granule)
		if ok && nextSlot != nil && nextSlot.Name == chunk.Name {
			// Rolling over chunk of the same name
			g.chunkSum -= nextSlot.Count
		}

		g.chunks.Enqueue(chunk)
		g.lastChunkName = chunkName

		// Reset the sliver buffer, but preserve its capacity
		cap := g.slivers.Capacity()
		g.slivers = &ring.Ring{}
		g.slivers.SetCapacity(cap)
	}

	g.slivers.Enqueue(struct{}{})
}

func (g *granularCounter) CountSlivers() int {
	g.RLock()
	defer g.RUnlock()

	return len(g.slivers.Values())
}

func (g *granularCounter) CountChunks() int {
	g.RLock()
	defer g.RUnlock()

	return g.chunkSum + len(g.slivers.Values())
}

type granule struct {
	Name  int
	Count int
}
