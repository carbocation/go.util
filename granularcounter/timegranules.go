package granularcounter

import (
	"time"
)

func Nanosecond() int {
	t := time.Now()

	return t.Nanosecond()
}

func Microsecond() int {
	t := time.Now()

	return t.Nanosecond() / 1000
}

func Millisecond() int {
	t := time.Now()

	return t.Nanosecond() / 1000000
}

// Qsec yields a quarter of a second
func Qsec() int {
	t := time.Now()

	return (t.Nanosecond() * 4 / 1000000000)
}

func Second() int {
	t := time.Now()

	return t.Second()
}

func Minute() int {
	t := time.Now()

	return t.Minute()
}

func Hour() int {
	t := time.Now()

	return t.Hour()
}

func Day() int {
	t := time.Now()

	return t.Day()
}

func Forever() int {
	return 1
}
