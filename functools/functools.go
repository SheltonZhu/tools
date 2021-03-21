package functools

type Iterable interface {
	Len() int
	Get(i int) Item
}

func Map(iterable Iterable, fn func(x Item) Item) (retVal ItemSlice) {
	for i := 0; i < iterable.Len(); i++ {
		retVal = append(retVal, fn(iterable.Get(i)))
	}
	return
}
func Reduce(iterable Iterable, fn func(x, y Item) Item) Item {
	var retVal = iterable.Get(0)
	for i := 1; i < iterable.Len(); i++ {
		retVal = fn(retVal, iterable.Get(i))
	}
	return retVal
}
func Filter(iterable Iterable, fn func(x Item) bool) (retVal ItemSlice) {
	for i := 0; i < iterable.Len(); i++ {
		if fn(iterable.Get(i)) {
			retVal = append(retVal, iterable.Get(i))
		}
	}
	return
}

type Item interface{}
type ItemSlice []Item

func (p ItemSlice) Len() int       { return len(p) }
func (p ItemSlice) Get(i int) Item { return p[i] }

type IntSlice []int

func (p IntSlice) Len() int       { return len(p) }
func (p IntSlice) Get(i int) Item { return p[i] }

func (p IntSlice) Map(fn func(item Item) Item) ItemSlice { return Map(p, fn) }
func (p IntSlice) Reduce(fn func(x, y Item) Item) Item   { return Reduce(p, fn) }
func (p IntSlice) Filter(fn func(x Item) bool) ItemSlice { return Filter(p, fn) }

type Float64Slice []float64

func (p Float64Slice) Len() int       { return len(p) }
func (p Float64Slice) Get(i int) Item { return p[i] }

func (p Float64Slice) Map(fn func(item Item) Item) ItemSlice { return Map(p, fn) }
func (p Float64Slice) Reduce(fn func(x, y Item) Item) Item   { return Reduce(p, fn) }
func (p Float64Slice) Filter(fn func(x Item) bool) ItemSlice { return Filter(p, fn) }

type StringSlice []string

func (p StringSlice) Len() int       { return len(p) }
func (p StringSlice) Get(i int) Item { return p[i] }

func (p StringSlice) Map(fn func(item Item) Item) ItemSlice { return Map(p, fn) }
func (p StringSlice) Reduce(fn func(x, y Item) Item) Item   { return Reduce(p, fn) }
func (p StringSlice) Filter(fn func(x Item) bool) ItemSlice { return Filter(p, fn) }

func MapInts(a []int, fn func(x Item) Item) ItemSlice         { return Map(IntSlice(a), fn) }
func MapFloat64s(a []float64, fn func(x Item) Item) ItemSlice { return Map(Float64Slice(a), fn) }
func MapStrings(a []string, fn func(x Item) Item) ItemSlice   { return Map(StringSlice(a), fn) }

func ReduceInts(a []int, fn func(x, y Item) Item) Item         { return Reduce(IntSlice(a), fn) }
func ReduceFloat64s(a []float64, fn func(x, y Item) Item) Item { return Reduce(Float64Slice(a), fn) }
func ReduceStrings(a []string, fn func(x, y Item) Item) Item   { return Reduce(StringSlice(a), fn) }

func FilterInts(a []int, fn func(x Item) bool) ItemSlice         { return Filter(IntSlice(a), fn) }
func FilterFloat64s(a []float64, fn func(x Item) bool) ItemSlice { return Filter(Float64Slice(a), fn) }
func FilterStrings(a []string, fn func(x Item) bool) ItemSlice   { return Filter(StringSlice(a), fn) }

func CompareSlice(a, b Iterable) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if a.Get(i) != b.Get(i) {
			return false
		}
	}
	return true
}
