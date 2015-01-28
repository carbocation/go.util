package granularcounter

import (
	"testing"
)

const (
	MaxPerMin = 240
	MaxPerDay = 120000
)

func TestMinuteDay(t *testing.T) {
	min := NewGranularCounter(minute, 240)
	hour := min.NewParent(hour, 60)
	day := hour.NewParent(day, 24)

	for i := 0; i < 2000; i++ {
		if min.SumChildren() >= MaxPerMin || day.SumChildren() >= MaxPerDay {
			//t.Log("Over quota")
			continue
		}

		min.Add(MakeCountable(1))
	}

	if actual := day.SumChildren(); actual > MaxPerMin {
		t.Errorf("Expected %d or fewer, got %d", MaxPerMin, actual)
	}

	t.Log(day.SumChildren())
}
