package functools

import (
	"strconv"
	"strings"
	"testing"
)

var (
	intData              = []int{1, 2, 3, 4, 5, 6}
	float64Data          = []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	stringData           = []string{"1", "2", "3", "4", "5", "6"}
	structData           = []newStruct{{"1", 1}, {"2", 2}, {"3", 3}}
	intMapExpected       = []int{2, 4, 6, 8, 10, 12}
	intMapFunc           = func(x Item) Item { return x.(int) * 2 }
	intReduceExpected    = 21
	intReduceFunc        = func(x, y Item) Item { return x.(int) + y.(int) }
	intFilterExpected    = []int{2, 4, 6}
	intFilterFunc        = func(x Item) bool { return x.(int)%2 == 0 }
	floatMapExpected     = []float64{2.2, 4.4, 6.6, 8.8, 11, 13.2}
	floatMapFunc         = func(x Item) Item { return x.(float64) * 2 }
	floatReduceExpected  = 23.1
	floatReduceFunc      = func(x, y Item) Item { return x.(float64) + y.(float64) }
	floatFilterExpected  = []float64{3.3, 4.4, 5.5, 6.6}
	floatFilterFunc      = func(x Item) bool { return x.(float64) >= 3 }
	stringMapExpected    = []string{"11", "22", "33", "44", "55", "66"}
	stringMapFunc        = func(x Item) Item { return strings.Repeat(x.(string), 2) }
	stringReduceExpected = "123456"
	StringReduceFunc     = func(x, y Item) Item { return x.(string) + y.(string) }
	stringFilterExpected = []string{"4"}
	stringFilterFunc     = func(x Item) bool { return strings.Contains(x.(string), "4") }
)

type newStruct struct {
	key string
	id  int
}
type newStructSlice []newStruct

func (p newStructSlice) Len() int {
	return len(p)
}
func (p newStructSlice) Get(i int) Item {
	return p[i]
}

func TestMap(t *testing.T) {
	expected := newStructSlice{newStruct{"10", 10}, newStruct{"20", 20}, newStruct{"30", 30}}
	actual := Map(newStructSlice(structData), func(x Item) Item {
		item := x.(newStruct)
		item.id = item.id * 10
		item.key = strconv.Itoa(item.id)
		return item
	})
	if !CompareSlice(actual, expected) {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, expected)
	}
}
func TestReduce(t *testing.T) {
	expected := newStruct{id: 6, key: "123"}
	actual := Reduce(newStructSlice(structData), func(x, y Item) Item {
		xx := x.(newStruct)
		yy := y.(newStruct)
		return newStruct{id: xx.id + yy.id, key: xx.key + yy.key}
	})
	if actual != expected {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, expected)
	}
}
func TestFilter(t *testing.T) {
	expected := newStructSlice{newStruct{"2", 2}, newStruct{"3", 3}}
	actual := Filter(newStructSlice(structData), func(x Item) bool {
		item := x.(newStruct)
		return item.id > 1
	})
	if !CompareSlice(actual, expected) {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, expected)
	}
}

func TestMapInts(t *testing.T) {
	actual := MapInts(intData, intMapFunc)
	if !CompareSlice(actual, IntSlice(intMapExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, intMapExpected)
	}
}
func TestMapFloat64s(t *testing.T) {
	actual := MapFloat64s(float64Data, floatMapFunc)
	if !CompareSlice(actual, Float64Slice(floatMapExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", float64Data, actual, floatMapExpected)
	}
}
func TestMapStrings(t *testing.T) {
	actual := MapStrings(stringData, stringMapFunc)
	if !CompareSlice(actual, StringSlice(stringMapExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", stringData, actual, stringMapExpected)
	}
}

func TestReduceInts(t *testing.T) {
	actual := ReduceInts(intData, intReduceFunc)
	if actual != intReduceExpected {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, intReduceExpected)
	}
}
func TestReduceFloat64s(t *testing.T) {
	actual := ReduceFloat64s(float64Data, floatReduceFunc)
	if actual != floatReduceExpected {
		t.Errorf("Map(%v) = %v; expected %v", float64Data, actual, floatReduceExpected)
	}
}
func TestReduceStrings(t *testing.T) {
	actual := ReduceStrings(stringData, StringReduceFunc)
	if actual != stringReduceExpected {
		t.Errorf("Map(%v) = %v; expected %v", stringData, actual, stringReduceExpected)
	}
}

func TestFilterInts(t *testing.T) {
	actual := FilterInts(intData, intFilterFunc)
	if !CompareSlice(actual, IntSlice(intFilterExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, intFilterExpected)
	}
}
func TestFilterFloat64s(t *testing.T) {
	actual := FilterFloat64s(float64Data, floatFilterFunc)
	if !CompareSlice(actual, Float64Slice(floatFilterExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", float64Data, actual, floatFilterExpected)
	}
}
func TestFilterStrings(t *testing.T) {
	actual := FilterStrings(stringData, stringFilterFunc)
	if !CompareSlice(actual, StringSlice(stringFilterExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", stringData, actual, stringFilterExpected)
	}
}

func TestIntSlice_Map(t *testing.T) {
	actual := IntSlice(intData).Map(intMapFunc)
	if !CompareSlice(actual, IntSlice(intMapExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, intMapExpected)
	}
}
func TestIntSlice_Reduce(t *testing.T) {
	actual := IntSlice(intData).Reduce(intReduceFunc)
	if actual != intReduceExpected {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, intReduceExpected)
	}
}
func TestIntSlice_Filter(t *testing.T) {
	actual := IntSlice(intData).Filter(intFilterFunc)
	if !CompareSlice(actual, IntSlice(intFilterExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", intData, actual, intFilterExpected)
	}
}

func TestFloat64Slice_Map(t *testing.T) {
	actual := Float64Slice(float64Data).Map(floatMapFunc)
	if !CompareSlice(actual, Float64Slice(floatMapExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", float64Data, actual, floatMapExpected)
	}
}
func TestFloat64Slice_Reduce(t *testing.T) {
	actual := Float64Slice(float64Data).Reduce(floatReduceFunc)
	if actual != floatReduceExpected {
		t.Errorf("Map(%v) = %v; expected %v", float64Data, actual, floatReduceExpected)
	}
}
func TestFloat64Slice_Filter(t *testing.T) {
	actual := Float64Slice(float64Data).Filter(floatFilterFunc)
	if !CompareSlice(actual, Float64Slice(floatFilterExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", float64Data, actual, floatFilterExpected)
	}
}

func TestStringSlice_Map(t *testing.T) {
	actual := StringSlice(stringData).Map(stringMapFunc)
	if !CompareSlice(actual, StringSlice(stringMapExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", stringData, actual, stringMapExpected)
	}
}
func TestStringSlice_Reduce(t *testing.T) {
	actual := StringSlice(stringData).Reduce(StringReduceFunc)
	if actual != stringReduceExpected {
		t.Errorf("Map(%v) = %v; expected %v", stringData, actual, stringReduceExpected)
	}
}
func TestStringSlice_Filter(t *testing.T) {
	actual := StringSlice(stringData).Filter(stringFilterFunc)
	if !CompareSlice(actual, StringSlice(stringFilterExpected)) {
		t.Errorf("Map(%v) = %v; expected %v", stringData, actual, stringFilterExpected)
	}
}
