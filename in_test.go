package goin

import "testing"

func Test_IntIn(t *testing.T) {
	ar := []int{1, 4, 5, 2, 90}
	value := 5
	test1 := IntIn(ar, value)
	if test1 != true {
		t.Error("IntIn should return true but get", test1)
	}
	test2 := IntIn(ar, 8)
	if test2 != false {
		t.Error("IntIn should return true but get", test2)
	}
}
func Test_Float64In(t *testing.T) {
	ar := []float64{1.2, 4.110, 5.98, 2.78, 9.0}
	value := 5.9
	test1 := Float64In(ar, value)
	if test1 != false {
		t.Error("IntIn should return true but get", test1)
	}
	test2 := Float64In(ar, 4.11)
	if test2 != true {
		t.Error("IntIn should return true but get", test2)
	}
}
func Test_StringIn(t *testing.T) {
	ar := []string{"apple", "orange", "pine", "lemon"}
	value := "appl"
	test1 := StringIn(ar, value)
	if test1 != false {
		t.Error("IntIn should return true but get", test1)
	}
	test2 := StringIn(ar, "orange")
	if test2 != true {
		t.Error("IntIn should return true but get", test2)
	}
}
