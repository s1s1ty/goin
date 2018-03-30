package goin

import (
	"sort"
)

// Interface is describing `goin` features
type Interface interface {
	IntIn(ar []int, value int) bool
	// Int64In(ar []int64, value int64) bool # need to implement
	// FloatIn(ar []int64, value int64) bool # need to implement
	// InterfaceIn(ar []int64, value int64) bool # need to implement
	Float64In(ar []float64, value float64) bool
	StringIn(ar []string, value string) bool
}

// IntIn returns boolean value based on int sequence contains value variable or not
func IntIn(ar []int, value int) bool {
	sort.Ints(ar)
	start := 0
	end := len(ar) - 1
	for start <= end {
		mid := (start + end) / 2
		if ar[mid] == value {
			return true
		} else if ar[mid] < value {
			start = mid + 1
		} else if ar[mid] > value {
			end = mid - 1
		}
	}
	return false
}

// Float64In returns boolean value based on Float64 sequence contains value variable or not
func Float64In(ar []float64, value float64) bool {
	sort.Float64s(ar)
	start := 0
	end := len(ar) - 1
	for start <= end {
		mid := (start + end) / 2
		if ar[mid] == value {
			return true
		} else if ar[mid] < value {
			start = mid + 1
		} else if ar[mid] > value {
			end = mid - 1
		}
	}
	return false
}

// StringIn returns boolean value based on string sequence contains value variable or not
func StringIn(ar []string, value string) bool {
	sort.Strings(ar)
	start := 0
	end := len(ar) - 1
	for start <= end {
		mid := (start + end) / 2
		if ar[mid] == value {
			return true
		} else if ar[mid] < value {
			start = mid + 1
		} else if ar[mid] > value {
			end = mid - 1
		}
	}
	return false
}
