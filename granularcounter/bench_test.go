package granularcounter

import (
	"testing"
)

func BenchmarkAddMinute(b *testing.B) {
	chunkCap, sliverCap := 60, 10000

	gc := NewGranularCounter(minute, chunkCap, sliverCap)

	for i := 0; i < b.N; i++ {
		gc.Add()
	}

	b.Log(b.N, "Events this sliver (minute):", gc.CountSlivers(), "Events across all", chunkCap, "chunks:", gc.CountChunks())
}

func BenchmarkAddSecond(b *testing.B) {
	chunkCap, sliverCap := 60, 10000

	gc := NewGranularCounter(second, chunkCap, sliverCap)

	for i := 0; i < b.N; i++ {
		gc.Add()
	}

	b.Log(b.N, "Events this sliver (second):", gc.CountSlivers(), "Events across all", chunkCap, "chunks:", gc.CountChunks())
}

func BenchmarkAddNanosecond(b *testing.B) {
	chunkCap, sliverCap := 1000000, 2

	gc := NewGranularCounter(nanosecond, chunkCap, sliverCap)

	for i := 0; i < b.N; i++ {
		gc.Add()
	}

	b.Log(b.N, "Events this sliver (nanosecond):", gc.CountSlivers(), "Events across all", chunkCap, "chunks:", gc.CountChunks())
}
