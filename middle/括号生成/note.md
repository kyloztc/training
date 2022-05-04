## 题目链接
[括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

## 思路
递归  
使用一个int类型left代表左侧括号的数量，使用一个int类型right代表右侧括号的数量  
每次递归可以添加一个左侧括号或者右侧括号，添加左侧括号的条件是left < n，添加右侧括号的条件是left > right，当left = n时递归结束，补全右侧括号即找到一个结果