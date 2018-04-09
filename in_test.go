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
		if output, err := Value(test.key).In(test.seq); output != test.expected {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("Test Failed: %v sequence and %v key inputed, %v expected, received: %v", test.seq, test.key, test.expected, output)
			}
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
		if output, err := Value(test.key).In(test.seq); output != test.expected {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("Test Failed: %v sequence and %v key inputed, %v expected, received: %v", test.seq, test.key, test.expected, output)
			}
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
		if output, err := Value(test.key).In(test.seq); output != test.expected {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("Test Failed: %v sequence and %v key inputed, %v expected, received: %v", test.seq, test.key, test.expected, output)
			}
		}

	}
}

func Test_InFloat64(t *testing.T) {

	seq := []float64{4.110, 5.98, 0.00000000000000001, 1.9999999999999996, 5.33333333333}
	seq = append(seq, 0.2*6)
	testCases := []struct {
		key      float64
		expected bool
	}{
		{
			key:      1.2,
			expected: true,
		},
		{
			key:      0.00000000000000000,
			expected: false,
		},
		{
			key:      5.33333333333,
			expected: true,
		},
	}
	for _, test := range testCases {
		if output, err := Value(test.key).In(seq); output != test.expected {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("Test Failed: %v sequence and %v key inputed, %v expected, received: %v", seq, test.key, test.expected, output)
			}
		}

	}
}

func Test_InFloat32(t *testing.T) {
	seq := []float32{1.2, 4.110, 5.98, 2.78, 1.9999999996999996, 0.00000000010000001}
	testCases := []struct {
		key      float32
		expected bool
	}{
		{
			key:      4.1,
			expected: false,
		},
		{
			key:      1.9999999999,
			expected: true,
		},
		{
			key:      0.00000000010000001, // support 10^9
			expected: true,
		},
	}
	for _, test := range testCases {
		if output, err := Value(test.key).In(seq); output != test.expected {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("Test Failed: %v sequence and %v key inputed, %v expected, received: %v", seq, test.key, test.expected, output)
			}
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
		if output, err := Value(test.key).In(seq); output != test.expected {
			if err != nil {
				t.Error(err)
			} else {
				t.Errorf("Test Failed: %v sequence and %v key inputed, %v expected, received: %v", seq, test.key, test.expected, output)
			}
		}

	}
}

func Test_In(t *testing.T) {
	ar1 := []string{"apple", "orange", "pine", "lemon"}
	test1, _ := Value("apple").In(ar1)
	if test1 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar1, "apple", true, test1)
	}

	ar2 := []int{500, 100, 10, 20, 80}
	test2, _ := Value(int(800)).In(ar2)
	if test2 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar2, 800, false, test2)
	}

	ar3 := []int32{500, 100, 10, 20, 80}
	test3, _ := Value(int32(100)).In(ar3)
	if test3 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar3, 100, true, test3)
	}

	ar4 := []int64{500, 100, 10, 20, 80}
	test4, _ := Value(int64(20)).In(ar4)
	if test4 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar4, 20, true, test4)
	}

	ar5 := []float64{50.10, 2.100, 1.570, 23.20, 80.0}
	test5, _ := Value(float64(1.5)).In(ar5)
	if test5 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar5, 1.5, false, test5)
	}

	ar6 := []float32{50.10, 2.100, 1.570, 23.20, 80.0}
	test6, _ := Value(float32(50)).In(ar6)
	if test6 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar6, 50, false, test6)
	}

	ar7 := [5]int{500, 100, 10, 20, 80}
	test7, _ := Value(int(800)).In(ar7)
	if test7 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar7, 800, false, test7)
	}

	ar8 := []interface{}{500, "100", 10, 20.1, 80}
	test8, _ := Value("100").In(ar8)
	if test8 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", ar8, "100", false, test8)
	}
}

func Test_InKey(t *testing.T) {

	mp1 := map[string]int{"junior": 1, "mid": 2, "senior": 3}
	test1, _ := Value("junior").InKey(mp1)
	if test1 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp1, "junior", true, test1)
	}

	mp2 := map[int]string{110: "shaon", 202: "munni", 390: "duli"}
	test2, _ := Value(10).InKey(mp2)
	if test2 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp2, 10, false, test2)
	}

	mp3 := map[string]string{"name": "shaon", "id": "102", "dept": "cse"}
	test3, _ := Value("fullname").InKey(mp3)
	if test3 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp3, "fullname", false, test3)
	}

	mp4 := map[int]float64{1: 222.90, 2: 345.87, 3: 500.78}
	test4, _ := Value(2).InKey(mp4)
	if test4 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp4, 2, true, test4)
	}

	mp5 := map[int]float32{1: 222.90, 2: 345.87, 3: 500.78}
	test5, _ := Value(2).InKey(mp5)
	if test5 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp5, 2, true, test5)
	}

	mp6 := map[int]float32{1: 222.90, 2: 345.87, 3: 500.78}
	test6, _ := Value(4).InKey(mp6)
	if test6 != false {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp6, 4, false, test6)
	}

	mp7 := map[interface{}]interface{}{1: 222.90, "2": 345.87, 3.1: 500.78}
	test7, _ := Value("2").InKey(mp7)
	if test7 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp7, "2", false, test7)
	}
	test7_1, _ := Value(3.1).InKey(mp7)
	if test7_1 != true {
		t.Errorf("Test Failed: %v sequence and %v key inputted, %v expected, received: %v", mp7, 3.1, false, test7_1)
	}
}
