package utils

import (
	"fmt"
	"time"
)

func TypeCheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int, int64:
			fmt.Printf("%+v is int.\n", v)
		case float64:
			fmt.Printf("%+v is float64.\n", v)
		case float32:
			fmt.Printf("%+v is float32.\n", v)
		case string:
			fmt.Printf("%+v is string.\n", v)
		case bool:
			fmt.Printf("%+v is bool.\n", v)
		case nil:
			fmt.Printf("%+v is nil.\n", v)
		default:
			fmt.Printf("%T is unknown.\n", v)
		}
	}
}

func RangeInts(start, end, step int) chan int {
	if start < 0 || end < 0 {
		panic("'start' must be positive integer.")
	}

	if step <= 0 {
		panic("'step' must be positive integer.")
	}

	ch := make(chan int)
	go func() {
		for tmp := start; tmp < end; tmp += step {
			ch <- tmp
		}
		close(ch)
	}()
	return ch
}

func RangeIntsSlice(start, end, step int) (slice []int) {
	ch := RangeInts(start, end, step)
	for i := range ch {
		slice = append(slice, i)
	}
	return
}

func DaysFromYear(year int) int {
	last := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)
	first := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	return int(last.Sub(first).Hours()/24 + 1)
}
