package functools

import (
	"fmt"
	"reflect"
)
// Example:
//
//	Map([]int{1, 2, 3}, func(x interface{}) interface{} {
//		return x.(int) * 2
//	})
//
// Return: [2, 4, 6]
func Map(iterable interface{}, fn func(x interface{}) interface{}) interface{} {
	switch reflect.TypeOf(iterable).Kind() {
	case reflect.Slice:
		var ret []interface{}
		iterValue := reflect.ValueOf(iterable)
		for i := 0; i < iterValue.Len(); i++ {
			ret = append(ret, fn(iterValue.Index(i).Interface()))
		}

		return ret

	default:
		panic(fmt.Sprintf("%v is not iterable.", iterable))
	}
}

// Example:
//
//	Reduce([]int{1, 2, 3}, func(x, y interface{}) interface{} {
//		return x.(int) + y.(int)
//	})
//
// Return: 6
func Reduce(iterable interface{}, fn func(x, y interface{}) interface{}) interface{} {
	switch reflect.TypeOf(iterable).Kind() {
	case reflect.Slice:
		iterValue := reflect.ValueOf(iterable)
		var ret = iterValue.Index(0).Interface()
		for i := 1; i < iterValue.Len(); i++ {
			ret = fn(ret, iterValue.Index(i).Interface())
		}
		return ret

	default:
		panic(fmt.Sprintf("%v is not iterable.", iterable))
	}
}

// Example:
//
//	FilterInt([]int{1, 2, 3}, func(x int) bool {
//		return x%2 == 0
//	})
//
// Return: [2]
func Filter(iterable interface{}, fn func(x interface{}) bool) interface{} {
	switch reflect.TypeOf(iterable).Kind() {
	case reflect.Slice:
		var ret []interface{}
		iterValue := reflect.ValueOf(iterable)
		for i := 0; i < iterValue.Len(); i++ {
			if fn(iterValue.Index(i).Interface()) {
				ret = append(ret, iterValue.Index(i).Interface())
			}
		}
		return ret

	default:
		panic(fmt.Sprintf("%v is not iterable.", iterable))
	}
}
