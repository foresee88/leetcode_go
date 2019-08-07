package main

import "fmt"

/*

17. Letter Combinations of a Phone Number
Medium

2334

315

Favorite

Share
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

*/

func letterCombinations(digits string) []string {
	m := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	strs,strs1 := make([]string, 0),make([]string, 0)

	if len(digits)<1{
		return []string{}
	}

	for _, s := range m[digits[0]] {
		strs = append(strs, string(s))
	}

	for i:=1; i<len(digits); i++ {

		for j:=0; j < len(strs);j++{

			for _, s := range m[digits[i]] {
				strs1 = append(strs1, strs[j]+string(s))
			}
		}
		strs = strs1
		strs1 = []string{}
	}
	return strs
}

func main() {
	fmt.Println(letterCombinations("234"))
}
