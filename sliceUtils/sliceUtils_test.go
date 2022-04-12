package sliceUtils

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	type testTemplate struct {
		input []int
		want  []int
	}

	testCases := []testTemplate{
		{input: []int{1, 3, 5}, want: []int{5, 3, 1}},
		{input: []int{10, 35, 105, 5}, want: []int{5, 105, 35, 10}},
		{input: []int{1}, want: []int{1}},
		{input: []int{51, 1, 5, 8, 9}, want: []int{9, 8, 5, 1, 51}},
	}

	for _, tc := range testCases {
		result := ReverseArray(tc.input)
		if !reflect.DeepEqual(tc.want, result) {
			t.Fatalf("Expected: %v, But Got: %v", tc.want, result)
		}
	}
}

func TestQuickSort(t *testing.T) {
	type testTemplate struct {
		input []int
		want  []int
	}

	testCases := []testTemplate{
		{input: []int{1, 51, 31, 156, 321}, want: []int{1, 31, 51, 156, 321}},
		{input: []int{2, 3}, want: []int{2, 3}},
		{input: []int{}, want: []int{}},
		{input: []int{15}, want: []int{15}},
	}

	for _, tc := range testCases {
		result := QuickSort(tc.want)

		if !reflect.DeepEqual(tc.want, result) {
			t.Fatalf("Expected: %v, But Got: %v", tc.want, result)
		}
	}

}
