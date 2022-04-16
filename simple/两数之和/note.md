## 题目链接
[两数之和](https://leetcode-cn.com/problems/two-sum/)
## 思路
使用`map[数值]数值下表列表`的方式存放原始数据（helper），然后遍历数组，使用`target-当前数值`得出另外一个数的大小，利用helper的key值快速的到下标列表，找到第一个与当前数值下标不一致的下标即为另一个数值下标。