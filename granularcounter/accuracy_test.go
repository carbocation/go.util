package granularcounter

import (
	"testing"
)

const (
	MaxPerMin = 240
	MaxPerDay = 120000
)

func TestMinuteDay(t *testing.T) {
	counter := NewGranularCounter(MaxPerMin)
	hour := counter.NewParent(Minute, 60)
	day := hour.NewParent(Hour, 24)

	for i := 0; i < 2000; i++ {
		if counter.SumChildren() >= MaxPerMin || day.SumChildren() >= MaxPerDay {
			//t.Log("Over quota")
			continue
		}

		counter.Add(MakeCountable(1))
	}

	if actual := day.SumChildren(); actual > MaxPerMin {
		t.Errorf("Expected %d or fewer, got %d", MaxPerMin, actual)
	}

	t.Log(day.SumChildren())
}

func TestSecondSecond(t *testing.T) {
	counter := NewGranularCounter(100)
	sec2 := counter.NewParent(Millisecond, 100)
	sec3 := sec2.NewParent(Minute, 100)

	for i := 0; i < 2000; i++ {
		counter.Add(MakeCountable(1))
	}

	t.Log("Largest:", sec3.SumChildren(), "Large:", sec2.SumChildren(), "Small:", counter.SumChildren())
}
