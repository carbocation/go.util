package granularcounter

import (
	"testing"
	"time"
)

func BenchmarkAddMinute(b *testing.B) {
	//nsec := NewGranularCounter(nanosecond, 10)
	//usec := NewGranularCounter(microsecond, 10)
	msec := NewGranularCounter(millisecond, 5000)
	sec := msec.NewParent(second, 1000)
	min := sec.NewParent(minute, 60)

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
