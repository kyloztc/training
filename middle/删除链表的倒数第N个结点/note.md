## 题目链接
[删除链表的倒数第N个节点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

## 思路
使用两个指针AB，指针A先走n个节点，然后AB同时走，当A为nil时，B就是要删除的节点  
因为要删除B节点需要B的上一个节点，这里要在原链表前加两个节点辅助，第一个节点用于B的前置节点的起始点，第二个节点为A的起始点