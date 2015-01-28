package granularcounter

import (
	"testing"
)

func BenchmarkForever(b *testing.B) {
	counter := NewGranularCounter(1)
	counter.NewParent(Forever, 1)

	for i := 0; i < b.N; i++ {
		counter.Add(MakeCountable(1))
	}

	b.Log("N=", b.N, "All:", counter.parent.SumChildren())
}

func BenchmarkQSec(b *testing.B) {
	counter := NewGranularCounter(1)
	counter.NewParent(Qsec, 4)
	counter.parent.NewParent(Second, 60)

	for i := 0; i < b.N; i++ {
		counter.Add(MakeCountable(1))
	}

	b.Log("N=", b.N, "All:", counter.parent.parent.SumChildren())
	for _, chunk := range counter.parent.Values() {
		val := chunk.(*countable)
		b.Log("Quarter second:", val.name, val.val)
	}
}

func BenchmarkAddMinute(b *testing.B) {
	counter := NewGranularCounter(1)
	msec := counter.NewParent(Millisecond, 10)
	all := msec.NewParent(Second, 60)

	for i := 0; i < b.N; i++ {
		counter.Add(MakeCountable(1))
	}

	b.Log("N=", b.N, "All:", all.SumChildren(), "in", all.Len(), "slots",
		"Last", msec.Len(), "milliseconds:", msec.SumChildren(),
		"Last", counter.Len(), "sub-millisecond:", counter.SumChildren(),
	)
}

func BenchmarkNoRollup(b *testing.B) {
	counter := NewGranularCounter(1)
	ns := counter.NewParent(Second, 10)

	for i := 0; i < b.N; i++ {
		counter.Add(MakeCountable(1))
	}

	b.Log("Ran", b.N, "All:", ns.SumChildren(), "in", ns.Len(), "slots")
}
