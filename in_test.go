package goin

import (
	"testing"
)

func Test_InInt(t *testing.T) {
	testCases := []struct {
		seq      []int
		key      int
		expected bool
	}{
		{
			seq:      []int{1, 4, 5, 2, 90},
			key:      5,
			expected: true,
		},
		{
			seq:      []int{1, 4, 5, 2, 90},
			key:      8,
			expected: false,
		},
	}
	for _, test := range testCases {
		if output := InInt(test.seq, test.key); output != test.expected {
			t.Error("Test Failed: {} sequence and {} key inputted, {} expected, recieved: {}", test.seq, test.key, test.expected, output)
		}
	}
}
func Test_InFloat64(t *testing.T) {
	seq := []float64{1.2, 4.110, 5.98, 2.78, 9.0}
	testCases := []struct {
		key      float64
		expected bool
	}{
		{
			key:      5.9,
			expected: false,
		},
		{
			key:      4.11,
			expected: true,
		},
	}
	for _, test := range testCases {
		if output := InFloat64(seq, test.key); output != test.expected {
			t.Error("Test Failed: {} sequence and {} key inputted, {} expected, recieved: {}", seq, test.key, test.expected, output)
		}
	}
}
func Test_InString(t *testing.T) {
	seq := []string{"apple", "orange", "pine", "lemon"}
	testCases := []struct {
		key      string
		expected bool
	}{
		{
			key:      "appl",
			expected: false,
		},
		{
			key:      "orange",
			expected: true,
		},
	}
	for _, test := range testCases {
		if output := InString(seq, test.key); output != test.expected {
			t.Error("Test Failed: {} sequence and {} key inputted, {} expected, recieved: {}", seq, test.key, test.expected, output)
		}
	}
}

func Test_In(t *testing.T) {
	ar1 := []string{"apple", "orange", "pine", "lemon"}
	test1 := In(ar1, "apple", "string")
	if test1 != true {
		t.Error("In should return true but get ", test1)
	}
	ar2 := []int{500, 100, 10, 20, 80}
	test2 := In(ar2, 800, "int")
	if test2 != false {
		t.Error("In should return true but get ", test2)
	}
	ar3 := []float64{50.10, 2.100, 1.570, 23.20, 80.0}
	test3 := In(ar3, 1.5, "float64")
	if test3 != false {
		t.Error("In should return true but get ", test3)
	}
}
