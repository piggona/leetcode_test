https://leetcode-cn.com/problems/palindrome-partitioning-ii/submissions/

## 题解

本题的关键是要提前算出字符串所有子序列的回文情况：
go
```
pstr := make([][]bool, len(s))
for i := 0; i < len(pstr); i++ {
    pstr[i] = make([]bool, len(s))
    for j := 0; j < len(s); j++ {
        pstr[i][j] = true
    }
}
for i := len(s) - 1; i >= 0; i-- {
    for j := i + 1; j < len(s); j++ {
        pstr[i][j] = s[i] == s[j] && pstr[i+1][j-1]
    }
}
```

然后再运用动态规划，规则为：

新加入的元素为c，那么它与它前面的字符串s的回文关系可能是：
1. c与s没有任何回文关系，那么s+c的分割数就是：`s的分割数+1`
2. c与s的某个后缀子串s'可以构成回文串，s''+s'=s,那么s+c的分割数为：`s''的分割数+1`

比较这两种情况的分割数，取最小即可得到答案