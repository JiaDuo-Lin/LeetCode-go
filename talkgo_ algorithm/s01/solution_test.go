package s01

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	answer := twoSum1([]int{1, 2, 4, 7, 11, 15}, 15)
	require.EqualValues(t, 4, answer[0])
	require.EqualValues(t, 11, answer[1])
	answer = twoSum1([]int{1, 2, 3, 7, 8, 10, 12}, 13)
	require.EqualValues(t, 1, answer[0])
	require.EqualValues(t, 12, answer[1])
	answer = twoSum1([]int{0, 2, 3, 10, 11, 44, 55, 66, 77}, 55)
	require.EqualValues(t, 0, answer[0])
	require.EqualValues(t, 55, answer[1])
	answer = twoSum1([]int{1, 2, 4, 7, 8, 11, 15}, 10)
	require.EqualValues(t, 2, answer[0])
	require.EqualValues(t, 8, answer[1])
}

func TestTwoSum2(t *testing.T) {
	answer := twoSum2([]int{1, 2, 4, 7, 11, 15}, 15)
	require.EqualValues(t, 4, answer[0])
	require.EqualValues(t, 11, answer[1])
	answer = twoSum2([]int{1, 2, 3, 7, 8, 10, 12}, 13)
	require.EqualValues(t, 1, answer[0])
	require.EqualValues(t, 12, answer[1])
	answer = twoSum2([]int{0, 2, 3, 10, 11, 44, 55, 66, 77}, 55)
	require.EqualValues(t, 0, answer[0])
	require.EqualValues(t, 55, answer[1])
	answer = twoSum1([]int{1, 2, 4, 7, 8, 11, 15}, 10)
	require.EqualValues(t, 2, answer[0])
	require.EqualValues(t, 8, answer[1])
}


func TestFindMedian(t *testing.T) {
	answer := findMedian([]int{1, 2, 4, 7, 11, 15})
	require.EqualValues(t, 5.5, answer)
	answer = findMedian([]int{1, 2, 3, 7, 8, 10, 12})
	require.EqualValues(t, 7., answer)
	answer = findMedian([]int{4, 5, 6, 1, 2})
	require.EqualValues(t, 4., answer)
	answer = findMedian([]int{20, 50, 60, 11, 15, 17})
	require.EqualValues(t, 18.5, answer)
}