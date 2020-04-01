package multi_test

import (
	multi "kata/multiply-array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multi_Input_Array_3_2_1_Shoul_Be_2_3_6(t *testing.T) {
	expected := []int{2, 3, 6}
	a := []int{3, 2, 1}

	actual := multi.Call(a)

	assert.Equal(t, expected, actual)
}
