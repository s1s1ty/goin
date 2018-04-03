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
			t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", test.seq, test.key, test.expected, output)
		}
	}
}
func Test_InInt32(t *testing.T) {
	testCases := []struct {
		seq      []int32
		key      int32
		expected bool
	}{
		{
			seq:      []int32{1, 4, 5, 2, 90},
			key:      90,
			expected: true,
		},
		{
			seq:      []int32{1, 4, 5, 2, 90},
			key:      8,
			expected: false,
		},
	}
	for _, test := range testCases {
		if output := InInt32(test.seq, test.key); output != test.expected {
			t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", test.seq, test.key, test.expected, output)
		}
	}
}
func Test_InInt64(t *testing.T) {
	testCases := []struct {
		seq      []int64
		key      int64
		expected bool
	}{
		{
			seq:      []int64{1, 4, 5, 2, 90},
			key:      2,
			expected: true,
		},
		{
			seq:      []int64{1, 4, 5, 2, 90},
			key:      8,
			expected: false,
		},
	}
	for _, test := range testCases {
		if output := InInt64(test.seq, test.key); output != test.expected {
			t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", test.seq, test.key, test.expected, output)
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
			t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", seq, test.key, test.expected, output)
		}
	}
}
func Test_InFloat32(t *testing.T) {
	seq := []float32{1.2, 4.110, 5.98, 2.78, 9.0}
	testCases := []struct {
		key      float32
		expected bool
	}{
		{
			key:      4.1,
			expected: false,
		},
		{
			key:      4.11,
			expected: true,
		},
	}
	for _, test := range testCases {
		if output := InFloat32(seq, test.key); output != test.expected {
			t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", seq, test.key, test.expected, output)
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
			t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", seq, test.key, test.expected, output)
		}
	}
}
func Test_In(t *testing.T) {
	ar1 := []string{"apple", "orange", "pine", "lemon"}
	test1 := In(ar1, "apple", "string")
	if test1 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar1, "apple", true, test1)
	}

	ar2 := []int{500, 100, 10, 20, 80}
	test2 := In(ar2, 800, "int")
	if test2 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar2, 800, false, test2)
	}

	ar3 := []int32{500, 100, 10, 20, 80}
	test3 := In(ar3, int32(100), "int32")
	if test3 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar3, 100, true, test3)
	}

	ar4 := []int64{500, 100, 10, 20, 80}
	test4 := In(ar4, int64(20), "int64")
	if test4 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar4, 20, true, test4)
	}

	ar5 := []float64{50.10, 2.100, 1.570, 23.20, 80.0}
	test5 := In(ar5, 1.5, "float64")
	if test5 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar5, 1.5, false, test5)
	}

	ar6 := []float32{50.10, 2.100, 1.570, 23.20, 80.0}
	test6 := In(ar6, float32(50), "float32")
	if test6 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar5, 50, false, test6)
	}

	mp1 := map[string]int{"junior": 1, "mid": 2, "senior": 3}
	test7 := In(mp1, "junior", "[string]int")
	if test7 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp1, "junior", true, test7)
	}

	mp2 := map[int]string{110: "shaon", 202: "munni", 390: "duli"}
	test8 := In(mp2, 10, "[int]string")
	if test8 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp2, 10, false, test8)
	}

	mp3 := map[string]string{"name": "shaon", "id": "102", "dept": "cse"}
	test9 := In(mp3, "fullname", "[string]string")
	if test8 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp3, "fullname", false, test9)
	}

	mp4 := map[int]float64{1: 222.90, 2: 345.87, 3: 500.78}
	test10 := In(mp4, 2, "[int]float64")
	if test10 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp4, 2, true, test10)
	}

	mp5 := map[int]float32{1: 222.90, 2: 345.87, 3: 500.78}
	test11 := In(mp5, 2, "[int]float32")
	if test11 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp5, 2, true, test11)
	}
}
