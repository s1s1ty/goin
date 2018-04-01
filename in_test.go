package goin

import "testing"

func Test_InInt(t *testing.T) {
	ar := []int{1, 4, 5, 2, 90}
	value := 5
	test1 := InInt(ar, value)
	if test1 != true {
		t.Error("InInt should return true but get", test1)
	}
	test2 := InInt(ar, 8)
	if test2 != false {
		t.Error("InInt should return true but get", test2)
	}
}
func Test_InFloat64(t *testing.T) {
	ar := []float64{1.2, 4.110, 5.98, 2.78, 9.0}
	value := 5.9
	test1 := InFloat64(ar, value)
	if test1 != false {
		t.Error("InFloat64 should return true but get", test1)
	}
	test2 := InFloat64(ar, 4.11)
	if test2 != true {
		t.Error("InFloat64 should return true but get", test2)
	}
}
func Test_InString(t *testing.T) {
	ar := []string{"apple", "orange", "pine", "lemon"}
	value := "appl"
	test1 := InString(ar, value)
	if test1 != false {
		t.Error("InString should return true but get", test1)
	}
	test2 := InString(ar, "orange")
	if test2 != true {
		t.Error("InString should return true but get", test2)
	}
}

func Test_In(t *testing.T) {
	ar := []string{"apple", "orange", "pine", "lemon"}
	test1 := In(ar, "apple", "string")
	if test1 != true {
		t.Error("In should return true but get ", test1)
	}
}
