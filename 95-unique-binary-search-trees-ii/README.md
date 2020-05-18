95. 不同的二叉搜索树 II

**难度：中等**

给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:

输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



### 思路

使用双指针方式，left、right分别指向生成数字的前后端。


输入为3时，递归示意图，

![image.png](https://pic.leetcode-cn.com/831b9f9a00fd36bb770549a34e693639be9021529b7c22855c361aa597d64b2e-image.png)



### 运行结果


执行用时 : 4 ms , 在所有 golang 提交中击败了 96.59% 的用户 

内存消耗 : 4.3 MB , 在所有 golang 提交中击败了 100.00% 的用户




