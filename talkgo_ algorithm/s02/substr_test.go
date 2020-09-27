package s02

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// 暴力解法：直接遍历寻找相等子串
// 时间O(NM)，空间O(1)
func Solution(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	// s1长度为0
	if l1 == 0 {
		return 0
	}
	// s2的长度不够
	if l2 == 0 || l2 < l1 {
		return -1
	}

	for i := 0; i <= l2 - l1; i++ {
		if s2[i : i + l1] == s1 {
			return i
		}
	}
	return -1
}

// KMP算法
// 时间O(M), 空间O(N)
func KMP(s1, s2 string) int {
	m:=len(s1)
	n:=len(s2)
	if m==0 {
		return 0
	}

	if n<m {
		return -1
	}
	next := computeNext(s1)

	q:=0
	for i:=0;i<n;i++ {
		for q>0 && s2[i]!=s1[q]{
			q=next[q-1]
		}
		if s2[i]==s1[q] {
			q++
		}
		if q == m {
			return i+1-m
		}
	}
	return -1

}

// 生成辅助数组
func computeNext(pattern string) []int {
	n:=len(pattern)
	next:=make([]int,n)
	k:=0
	for i:=1; i<n; i++ {
		for k>0 && pattern[k]!=pattern[i] {
			k=next[k-1]
		}
		if pattern[k]==pattern[i] {
			k++
		}
		next[i]=k
	}
	return next
}

// 测试
type testNode1 struct {
	str1     string
	str2     string
	expected int
}

var data1 = []testNode1{
	{"abc", "banana", -1},
	{"abc", "cba", -1},
	{"hello", "hello world", 0},
	{"nice", "nicnicnihenicenoce", 10},
	{"abcba", "sjkfabgbaabcbasalkd", 9},
}

func TestAllInString3(t *testing.T) {
	for _, v := range data1 {
		answer := Solution(v.str1, v.str2)
		require.EqualValues(t, v.expected, answer)
	}
}

func TestAllInString4(t *testing.T) {
	for _, v := range data1 {
		answer := KMP(v.str1, v.str2)
		require.EqualValues(t, v.expected, answer)
	}
}
