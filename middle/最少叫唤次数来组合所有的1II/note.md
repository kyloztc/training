## 题目思路
[最少交换次数来组合所有的 1 II](https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together-ii/)

## 思路
1. 滑动窗口，窗口的大小就是数组中1的数量
2. 需要注意的是数组是循环的，需要处理一下循环的情况
3. 可以使用前缀和来快速获得窗口内0的数量