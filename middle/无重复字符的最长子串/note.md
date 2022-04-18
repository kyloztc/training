## 题目链接
[无重复字符串的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

## 思路
用滑动窗口，要注意
1. helper中下标如果比left小（说明不在窗口内），则left不更新
2. right每次循环都向右移动，并且更新helper下标
