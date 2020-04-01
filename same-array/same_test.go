package same_test

import (
	"kata/same-array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Comp_Input_Array_A_And_B_Should_Be_True(t *testing.T) {
	expected := true
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	actual := same.Comp(a, b)

	assert.Equal(t, expected, actual)
}

func Test_Comp_Input_Array_A_And_B_Should_Be_False(t *testing.T) {
	expected := false
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	actual := same.Comp(a, b)

	assert.Equal(t, expected, actual)
}

func Test_Comp_Input_Array_A_nil_And_B_Should_Be_False(t *testing.T) {
	expected := false
	var a []int
	b := []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	actual := same.Comp(a, b)

	assert.Equal(t, expected, actual)
}

func Test_Comp_Input_Array_A_And_B_Swap_Index_Should_Be_True(t *testing.T) {
	expected := true
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{11 * 11, 121 * 121, 19 * 19, 144 * 144, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	actual := same.Comp(a, b)

	assert.Equal(t, expected, actual)
}
