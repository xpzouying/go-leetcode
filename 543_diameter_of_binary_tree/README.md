543. 二叉树的直径


难度：简单


给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \     
      4   5    
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


### 解题思路

1. 对于二叉树的直径长度，有可能不会穿过根节点。
2. 由于只需要给出最大值，所以取一个值max记录一下即可。

### 操作流程

使用递归的方式，递归遍历左右子节点，相加获得直径，与max进行比较：

1. 如果大于max，则更新max值；
2. 否则，忽略；


### 执行结果

执行结果：
通过 显示详情 执行用时 : 4 ms , 在所有 golang 提交中击败了 95.71% 的用户 

内存消耗 :
4.5 MB , 在所有 golang 提交中击败了 91.67% 的用户