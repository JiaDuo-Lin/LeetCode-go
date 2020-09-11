# 树
[TOC]

## 层次遍历

层次遍历主要借助队列来完成



#### 题目1：[637. 二叉树的层平均值](https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/)

**题目描述：**给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

**思路：**直接使用队列就好，这里有个小技巧，每次遍历一层的时候先把该层的数目记下`n=len(queue)`，用标记记下来n个数据为该层的元素，避免后面插入新的数据后造成混乱，也可以节省辅助数组的空间。

题解：

``` go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func averageOfLevels(root *TreeNode) (res []float64) {
    if root == nil {
        return nil
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    n, sum := 0, 0.0
    for len(queue) > 0 {
        n = len(queue)
        sum = 0.
        for i:=0; i<n; i++ {
            sum += float64(queue[0].Val)
            if queue[0].Left != nil {
                queue =append(queue, queue[0].Left)
            } 
            if queue[0].Right != nil {
                queue =append(queue, queue[0].Right)
            }
            queue = queue[1:]
        }
        res = append(res, sum/float64(n))
    }
    return 
}
```



