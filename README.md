LeetCode
========

## 仓库目的

本仓库用于记录和归档 LeetCode 题目的 Go 语言题解，主要目的是练习算法与数据结构、保留不同时期的解题思路，并按题号整理，方便回顾、检索和分享。

题解统一放在 [`Algorithms/`](./Algorithms) 目录，均使用 `package main`。多数题解文件带有 `main` 函数用于演示示例；链表等公共数据结构与辅助函数集中在 [`Algorithms/common.go`](./Algorithms/common.go)。带 `main` 的题解可通过 `go run Algorithms/<文件名>.go Algorithms/common.go` 单独运行。

## 更新规范

1. **目录与命名**：题解统一放在 `Algorithms/` 目录，文件名为 `题号_英文标题.go`（如 `15_3Sum.go`）；标题取自 LeetCode 原题，单词之间用下划线连接。
2. **文件结构**：保持 `package main` 形式，文件开头保留原题描述注释，随后给出解法函数；可复用的数据结构或辅助函数放入 `common.go`，避免在多个题解文件中重复定义。
3. **README 表格**：新增题解后，在下方表格中按题号升序补充对应行：`Title` 链接到 LeetCode 原题页面，`Solution` 填写 `[go](./Algorithms/文件名.go)`，`Difficulty` 按原题填写 Easy/Medium/Hard；尚未实现的题目将 `Solution` 留空。
4. **提交**：每次提交尽量只包含一道或少量相关题目的改动，提交信息简要说明本次更新的题号或内容（如 `22`、`readme`）。

### LeetCode Algorithm

| # | Title | Solution | Difficulty |
|---| ----- | -------- | ---------- |
|1|[Two Sum](https://leetcode.com/problems/two-sum/)| [go](./Algorithms/1_Two_Sum.go)|Easy|
|2|[Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)| [go](./Algorithms/2_Add_Two_Numbers.go)|Medium|
|3|[Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)| [go](./Algorithms/3_Longest_Substring_Without_Repeating_Characters.go)|Medium|
|4|[Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)| [go](./Algorithms/4_Median_of_Two_Sorted_Arrays.go)|Hard|
|5|[Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)| [go](./Algorithms/5_Longest_Palindromic_Substring.go)|Medium|
|6|[ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)| [go](./Algorithms/6_ZigZag_Conversion.go)|Easy|
|7|[Reverse Integer](https://leetcode.com/problems/reverse-integer/)| [go](./Algorithms/7_Reverse_Integer.go)|Easy|
|8|[String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/)| [go](./Algorithms/8_String_to_Integer.go)|Easy|
|9|[Palindrome Number](https://leetcode.com/problems/palindrome-number/)| [go](./Algorithms/9_Palindrome_Number.go)|Easy|
|10|[Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)| [go](./Algorithms/10_Regular_Expression_Matching.go)|Hard|
|11|[Container With Most Water](https://leetcode.com/problems/container-with-most-water/)| [go](./Algorithms/11_Container_With_Most_Water.go)|Medium|
|12|[Integer to Roman](https://leetcode.com/problems/integer-to-roman/)| [go](./Algorithms/12_Integer_to_Roman.go)|Medium|
|13|[Roman to Integer](https://leetcode.com/problems/roman-to-integer/)| [go](./Algorithms/13_Roman_to_Integer.go)|Easy|
|14|[Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)| [go](./Algorithms/14_Longest_Common_Prefix.go)|Easy|
|15|[3Sum](https://leetcode.com/problems/3sum/)| [go](./Algorithms/15_3Sum.go)|Medium|
|16|[3Sum Closest](https://leetcode.com/problems/3sum-closest/)| [go](./Algorithms/16_3Sum_Closest.go)|Medium|
|17|[Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)| [go](./Algorithms/17_Letter_Combinations_of_a_Phone_Number.go)|Medium|
|18|[4Sum](https://leetcode.com/problems/4sum/)| [go](./Algorithms/18_4Sum.go)|Medium|
|19|[Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)| [go](./Algorithms/19_Remove_Nth_Node_From_End_of_List.go)|Easy|
|20|[Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)| [go](./Algorithms/20_Valid_Parentheses.go)|Easy|
|21|[Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)| [go](./Algorithms/21_Merge_Two_Sorted_Lists.go)|Easy|
|22|[Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)| [go](./Algorithms/22_Generate_Parentheses.go)|Medium|
|23|[Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)| [go](./Algorithms/23_Merge_k_Sorted_Lists.go)|Hard|
|24|[Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/)| [go](./Algorithms/24_Swap_Nodes_in_Pairs.go)|Medium|
|25|[Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/)| [go](./Algorithms/25_Reverse_Nodes_in_k-Group.go)|Hard|
|26|[Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)| |Easy|
|27|[Remove Element](https://leetcode.com/problems/remove-element/)| |Easy|
|28|[Implement strStr()](https://leetcode.com/problems/implement-strstr/)| |Easy|
|29|[Divide Two Integers](https://leetcode.com/problems/divide-two-integers/)| |Medium|
|30|[Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)| |Hard|
|31|[Next Permutation](https://leetcode.com/problems/next-permutation/)| |Medium|
|32|[Longest Valid Parentheses](https://leetcode.com/problems/longest-valid-parentheses/)| |Hard|
|33|[Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)| [go](./Algorithms/33_Search_in_Rotated_Sorted_Array.go)|Hard|
|34|[Search for a Range](https://leetcode.com/problems/search-for-a-range/)| |Medium|
|35|[Search Insert Position](https://leetcode.com/problems/search-insert-position/)| |Medium|
|36|[Valid Sudoku](https://leetcode.com/problems/valid-sudoku/)| |Easy|
|37|[Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)| |Hard|
|38|[Count and Say](https://leetcode.com/problems/count-and-say/)| |Easy|
|39|[Combination Sum](https://leetcode.com/problems/combination-sum/)| |Medium|
|40|[Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)| |Medium|
|41|[First Missing Positive](https://leetcode.com/problems/first-missing-positive/)| |Hard|
|42|[Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)| |Hard|
|43|[Multiply Strings](https://leetcode.com/problems/multiply-strings/)| |Medium|
|44|[Wildcard Matching](https://leetcode.com/problems/wildcard-matching/)| |Hard|
|45|[Jump Game II](https://leetcode.com/problems/jump-game-ii/)| |Hard|
|46|[Permutations](https://leetcode.com/problems/permutations/)| |Medium|
|47|[Permutations II](https://leetcode.com/problems/permutations-ii/)| |Hard|
|48|[Rotate Image](https://leetcode.com/problems/rotate-image/)| |Medium|
|49|[Group Anagrams](https://leetcode.com/problems/anagrams/)| |Medium|
|50|[Pow(x, n)](https://leetcode.com/problems/powx-n/)| |Medium|