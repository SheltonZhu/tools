package functools

import (
	"reflect"
	"strings"
	"testing"
)

func CompareSlice(a, b interface{}) bool {
	if reflect.TypeOf(a).Kind() == reflect.Slice && reflect.TypeOf(b).Kind() == reflect.Slice {
		iterA := reflect.ValueOf(a)
		iterB := reflect.ValueOf(b)
		if iterA.Len() != iterB.Len() {
			return false
		}
		for i := 0; i < iterA.Len(); i++ {
			if iterA.Index(i).Interface() != iterB.Index(i).Interface() {
				return false
			}
		}
	}
	return true
}

func TestMap(t *testing.T) {
	var testData = [][]interface{}{
		{[]int{1, 2, 3, 4, 5, 6}, func(x interface{}) interface{} {
			return x.(int) * 2
		}, []int{2, 4, 6, 8, 10, 12}},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}, func(x interface{}) interface{} {
			return x.(float64) * 2
		}, []float64{2.2, 4.4, 6.6, 8.8, 11, 13.2}},
		{[]string{"1", "2", "3", "4", "5", "6"}, func(x interface{}) interface{} {
			return strings.Repeat(x.(string), 2)
		}, []string{"11", "22", "33", "44", "55", "66"}},
	}

	for _, row := range testData {
		actual := Map(row[0], row[1].(func(x interface{}) interface{}))
		func(expected interface{}) {
			if !CompareSlice(actual, expected) {
				t.Errorf("Map(%v) = %v; expected %v", row[0], actual, expected)
			}
		}(row[2])

	}

}
func TestReduce(t *testing.T) {
	var testData = [][]interface{}{
		{[]int{1, 2, 3, 4, 5, 6}, func(x, y interface{}) interface{} {
			return x.(int) + y.(int)
		}, 21},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}, func(x, y interface{}) interface{} {
			return x.(float64) + y.(float64)
		}, 23.1},
		{[]string{"1", "2", "3", "4", "5", "6"}, func(x, y interface{}) interface{} {
			return x.(string) + y.(string)
		}, "123456"},
	}

	for _, row := range testData {
		actual := Reduce(row[0], row[1].(func(x, y interface{}) interface{}))
		expected := row[2]
		if actual != expected {
			t.Errorf("Map(%v) = %v; expected %v", row[0], actual, expected)
		}
	}
}

func TestFilter(t *testing.T) {
	var testData = [][]interface{}{
		{[]int{1, 2, 3, 4, 5, 6}, func(x interface{}) bool {
			return x.(int)%2 == 0
		}, []int{2, 4, 6}},
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}, func(x interface{}) bool {
			return x.(float64) >= 3
		}, []float64{3.3, 4.4, 5.5, 6.6}},
		{[]string{"1", "2", "3", "4", "5", "6"}, func(x interface{}) bool {
			return strings.Contains(x.(string), "4")
		}, []string{"4"}},
	}

	for _, row := range testData {
		actual := Filter(row[0], row[1].(func(x interface{}) bool))
		func(expected interface{}) {
			if !CompareSlice(actual, expected) {
				t.Errorf("Map(%v) = %v; expected %v", row[0], actual, expected)
			}
		}(row[2])
	}
}
