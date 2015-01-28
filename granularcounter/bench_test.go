package granularcounter

import (
	"testing"
)

func BenchmarkAddMinute(b *testing.B) {
	//nsec := NewGranularCounter(nanosecond, 10)
	//usec := NewGranularCounter(microsecond, 10)
	msec := NewGranularCounter(Millisecond, 5000)
	sec := msec.NewParent(Second, 1000)
	min := sec.NewParent(Minute, 60)

	for i := 0; i < b.N; i++ {
		msec.Add(MakeCountable(1))
	}

	b.Log("Ran", b.N, "Last min:", min.SumChildren(), min.Len(),
		"Last second:", sec.SumChildren(), "in", sec.Len(), "slots",
		"Last millisecond:", msec.SumChildren(), "in", msec.Len(), "slots",
		//"Last usecond:", usec.SumChildren(), usec.Len(),
	)

	/*
		for _, x := range min.Values() {
			if val, ok := x.(countable); ok {
				b.Log("Second:", val.name, "Count:", val.val)
			}
		}

		b.Log("And sub-second:", min.SumChildren()-min.Sum())
	*/
}
