package sort

import "sort"

// Float64s is just an adapter to use at init.
func Float64s(s []float64) []float64 {
	sort.Float64s(s)
	return s
}
