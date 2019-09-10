# 139 Word Break


### 题目描述

给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


### 我的推到过程



输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。


leetcode

size_s := 8


mem := make([]int, 9)

mem[0] = true


for i := 1; i < 9; i++:


1. i == 1
    preMem := mem[0] ==> true

    for j := 0; j < 1; j++ {

        1.1 suffix := s[0:1] ==> l

    }

mem[1] = false

2. i == 2:
    2.1  j==0
    preMem := mem[0] ==> true
    suffix := s[0:2] ==> le,

    2.2  j==1
    preMem := mem[1] ==> false

mem[2] = false

3. i == 3:
    3.1  j==0
    preMem := mem[0] ==> true
    suffix := s[0:3] ==> lee,

    3.2  j==1:
    preMem := mem[1] ==> false

    3.3  j==2:
    preMem := mem[2] ==> false

mem[3] = false

4. i == 4:
    4.1 j==0
    suffix := s[0:4] ==> leet
    ==> find_dict ==> true

mem[4] = true

5. i == 5:
    5.1
    ... 
    5.5 j==4:
    preMem := mem[4] ==> true
    suffix := s[4:] ==> code ==> true


...

6. i == 9:

    6.1 ...

    6.5 j==4:
    preMem := mem[4] ==> true
    suffix := s[4:] ==> code ==> true

完结