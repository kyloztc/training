## 题目连接
[和至少为K的最短子数组](https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/)

## 思路
1. 因为是求连续子数组的和，应该想到用前缀和
2. 使用前缀和后暴力也可以解决，但是时间复杂度会比较高
3. 这个时候要结合题意排除一些不可能的结果，可以使用一个helper队列记录还需要考虑那些index，按照遍历顺序把index顺序放入helper中，helper中剔除不可能index的策略如下
4. 如果preSum[curIndx] - preSum[helper[0]] >= k，则可以把helper[0]弹出，因为如果preSum[curIndx] - preSum[helper[1]] >= k，后者的子数组长度肯定比前者短。循环弹出helper队头，直至不再满足>=k，在过程中记录最小值
5. 除了队头的index可以弹出，队尾的也可以弹出，如果preSum[helper[len(helper)-1]] >= curSum，那么curSum后面的nextCurSum减去curSum肯定会比nextCurSum-preSum[helper[len(helper)-1]]大，如果preSum[helper[len(helper)-1]]能满足题意，则curSum也可以满足题意，并且长度更小，因此可以把队尾弹出。
6. 完成以上两步操作之后将curSum的index放入helper队尾。