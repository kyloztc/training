## 题目链接
[最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)

## 思路
构建前缀树  
在遍历前缀树获取最长前缀时有两种情况表示最长前缀已经结束：
1. 子节点不止有一个，意味着出现了分叉。
2. 有字符串在当前节点结束  

第一种情况比较简单，第二种情况需要借助一个flag标识当前节点是否有单词结束，注意不要用false覆盖true（有可能一个较长的单词把较短的单词覆盖了） 
