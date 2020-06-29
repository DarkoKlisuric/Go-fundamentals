package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 43},
		{[]int{32, 22, 30}, 83},
	}

	for _, v := range tests {
		x := mySum(v.data...)

		if x != v.answer {
			t.Errorf("Expected %d got %d", mySum(v.data...), v.answer)
		}
	}

	/*
	--- FAIL: TestMySum (0.00s)
		main_test.go:21: Expected 42 got 43
		main_test.go:21: Expected 84 got 83
	FAIL
	exit status 1
	FAIL	Go-fundamentals/Testing/01-test/02-table-test	0.001s
	*/
}
