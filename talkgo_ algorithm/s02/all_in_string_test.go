package s02

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 用str1的字符，去str2中找有没有相等，如果没有相等的就直接返回false
// 时间复杂度O(N*M) 空间复杂度O(1)
func Solution1(str1, str2 string) bool {
	for _, v := range str1 {
		flag := false
		for _, k := range str2 {
			if v == k {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

// 给str2建立一个map便签，用于表示str2中的字符存在状况
// 遍历str1，利用map来判断该字符是否存在于str2中
// 时间复杂度O(N) 空间复杂度O(N)
func Solution2(str1, str2 string) bool {
	m := make(map[rune]bool)

	for _, v := range str2 {
		m[v] = true
	}

	for _, v := range str1 {
		if _, ok := m[v]; !ok {
			return false
		}
	}
	return true
}

type testNode struct {
	str1     string
	str2     string
	expected bool
}

var data = []testNode{
	{"abc", "banana", false},
	{"abc", "cba", true},
	{"hello", "hello world", true},
	{"test", "tesla", true},
	{"there is a kitty", "there is a cat", false},
}

func TestAllInString1(t *testing.T) {
	for _, v := range data {
		answer := Solution1(v.str1, v.str2)
		require.EqualValues(t, v.expected, answer)
	}
}

func TestAllInString2(t *testing.T) {
	for _, v := range data {
		answer := Solution2(v.str1, v.str2)
		require.EqualValues(t, v.expected, answer)
	}
}
