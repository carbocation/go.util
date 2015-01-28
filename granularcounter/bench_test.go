package granularcounter

import (
	"testing"
)

func BenchmarkAddMinute(b *testing.B) {
	//b.Log("Starting again for", b.N)

	//nsec := NewGranularCounter(nanosecond, 10)
	//usec := NewGranularCounter(microsecond, 10)
	msec := NewGranularCounter(millisecond, 5000)
	sec := msec.NewParent(second, 1000)
	min := sec.NewParent(minute, 60)

	for i := 0; i < b.N; i++ {
		msec.Add(1)
	}

	b.Log(b.N, "events. Total (minute):", min.SumChildren(), min.Len(),
		"Last second:", sec.SumChildren(), sec.Len(),
		"Last millisecond:", msec.SumChildren(), msec.Len(),
		//"Last usecond:", usec.SumChildren(), usec.Len(),
	)
}
