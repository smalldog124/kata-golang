package sum_test

import (
	"kata/sum-array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Call_Input_Array_0_1_3_6_10_Should_Be_20_20_19_16_10_0(t *testing.T) {
	expected := []int{20, 20, 19, 16, 10, 0}
	input := []int{0, 1, 3, 6, 10}

	actual := sum.Call(input)

	assert.Equal(t, actual, expected)
}

func Test_Call_Input_Array_1_2_3_4_5_6_Should_Be_21_20_18_15_11_6_0(t *testing.T) {
	expected := []int{21, 20, 18, 15, 11, 6, 0}
	input := []int{1, 2, 3, 4, 5, 6}

	actual := sum.Call(input)

	assert.Equal(t, actual, expected)
}
