package goin

import (
	"log"
	"reflect"
	"sort"
)

// InInt returns boolean value based on Int sequence contains value variable or not
func InInt(ar []int, value int) bool {
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

// InInt32 returns boolean value based on Int32 sequence contains value variable or not
func InInt32(ar []int32, value int32) bool {
	sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
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

// InInt64 returns boolean value based on Int64 sequence contains value variable or not
func InInt64(ar []int64, value int64) bool {
	sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
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

// InFloat32 returns boolean value based on Float32 sequence contains value variable or not
func InFloat32(ar []float32, value float32) bool {
	sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
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

// InFloat64 returns boolean value based on Float64 sequence contains value variable or not
func InFloat64(ar []float64, value float64) bool {
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

// InString returns boolean value based on string sequence contains value variable or not
func InString(ar []string, value string) bool {
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

// In returns boolean based on interface
func In(ar interface{}, value interface{}, typ string) bool {
	switch typ {
	case "int":
		seq, ok := ar.([]int)
		if ok != true {
			log.Fatal("Convert type is int but sequence type is ", reflect.TypeOf(ar))
		}
		val, ok := value.(int)
		if ok != true {
			log.Fatal("Convert type is int but value type is ", reflect.TypeOf(val))
		}
		return InInt(seq, val)
	case "int32":
		seq, ok := ar.([]int32)
		if ok != true {
			log.Fatal("Convert type is int32 but sequence type is ", reflect.TypeOf(ar))
		}
		val, ok := value.(int32)
		if ok != true {
			log.Fatal("Convert type is int32 but value type is ", reflect.TypeOf(val))
		}
		return InInt32(seq, val)
	case "int64":
		seq, ok := ar.([]int64)
		if ok != true {
			log.Fatal("Convert type is int64 but sequence type is ", reflect.TypeOf(ar))
		}
		val, ok := value.(int64)
		if ok != true {
			log.Fatal("Convert type is int64 but value type is ", reflect.TypeOf(val))
		}
		return InInt64(seq, val)
	case "float32":
		seq, ok := ar.([]float32)
		if ok != true {
			log.Fatal("Convert type is float32 but sequence type is ", reflect.TypeOf(ar))
		}
		val, ok := value.(float32)
		if ok != true {
			log.Fatal("Convert type is float32 but value type is ", reflect.TypeOf(val))
		}
		return InFloat32(seq, val)
	case "float64":
		seq, ok := ar.([]float64)
		if ok != true {
			log.Fatal("Convert type is float64 but sequence type is ", reflect.TypeOf(ar))
		}
		val, ok := value.(float64)
		if ok != true {
			log.Fatal("Convert type is float64 but value type is ", reflect.TypeOf(val))
		}
		return InFloat64(seq, val)
	case "string":
		seq, ok := ar.([]string)
		if ok != true {
			log.Fatal("Convert type is string but sequence type is ", reflect.TypeOf(ar))
		}
		val, ok := value.(string)
		if ok != true {
			log.Fatal("Convert type is string but value type is ", reflect.TypeOf(val))
		}
		return InString(seq, val)
	case "[string]int":
		seq, ok := ar.(map[string]int)
		if !ok {
			log.Fatal("Convert type is map[string]int but sequence type is ", reflect.TypeOf(ar))
		}
		key, ok := value.(string)
		if !ok {
			log.Fatal("Convert type is string but value type is ", reflect.TypeOf(key))
		}
		if _, exist := seq[key]; exist {
			return true
		}
	case "[string]string":
		seq, ok := ar.(map[string]string)
		if !ok {
			log.Fatal("Convert type is map[string]string but sequence type is ", reflect.TypeOf(ar))
		}
		key, ok := value.(string)
		if !ok {
			log.Fatal("Convert type is string but value type is ", reflect.TypeOf(key))
		}
		if _, exist := seq[key]; exist {
			return true
		}
	case "[int]string":
		seq, ok := ar.(map[int]string)
		if !ok {
			log.Fatal("Convert type is map[int]string but sequence type is ", reflect.TypeOf(ar))
		}
		key, ok := value.(int)
		if !ok {
			log.Fatal("Convert type is int but value type is ", reflect.TypeOf(key))
		}
		if _, exist := seq[key]; exist {
			return true
		}
	case "[string]float32":
		seq, ok := ar.(map[string]float32)
		if !ok {
			log.Fatal("Convert type is map[int]string but sequence type is ", reflect.TypeOf(ar))
		}
		key, ok := value.(string)
		if !ok {
			log.Fatal("Convert type is string but value type is ", reflect.TypeOf(key))
		}
		if _, exist := seq[key]; exist {
			return true
		}
	case "[string]float64":
		seq, ok := ar.(map[string]float64)
		if !ok {
			log.Fatal("Convert type is map[string]float64 but sequence type is ", reflect.TypeOf(ar))
		}
		key, ok := value.(string)
		if !ok {
			log.Fatal("Convert type is string but value type is ", reflect.TypeOf(key))
		}
		if _, exist := seq[key]; exist {
			return true
		}
	case "[int]float32":
		seq, ok := ar.(map[int]float32)
		if !ok {
			log.Fatal("Convert type is map[int]float32 but sequence type is ", reflect.TypeOf(ar))
		}
		key, ok := value.(int)
		if !ok {
			log.Fatal("Convert type is int but value type is ", reflect.TypeOf(key))
		}
		if _, exist := seq[key]; exist {
			return true
		}
	case "[int]float64":
		seq, ok := ar.(map[int]float64)
		if !ok {
			log.Fatal("Convert type is map[int]float64 but sequence type is ", reflect.TypeOf(ar))
		}
		key, ok := value.(int)
		if !ok {
			log.Fatal("Convert type is int but value type is ", reflect.TypeOf(key))
		}
		if _, exist := seq[key]; exist {
			return true
		}
	default:
		log.Fatal("Unknown type")
	}
	return false
}
