/*
22. Generate Parentheses
Medium

12381

481

Add to List

Share
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.



Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]


Constraints:

1 <= n <= 8
Accepted
1,027,264
Submissions
1,470,142
*/

package main

import "time"

func generateParenthesis(n int) []string{
	var astr []string
	getAll("(", 1, n*2, &astr)
	return astr
}

func getAll(str string, cur int, dep int, astr *[]string) {
	if cur == dep {
		if valid(str){
			*astr = append(*astr, str)
		}
		return
	} else {
		getAll(str +"(", cur+1, dep, astr)
		getAll(str +")", cur+1, dep, astr)
	}
}

func valid(str string) bool{
	balance := 0
	for _, s := range str{
		if s == '(' {
			balance++
		} else {
			balance--
		}
		if balance <0 {
			return false
		}
	}
	return balance == 0
}

func generateParenthesis1(n int) []string{
	var astr []string
	getAll1("(", 1, n*2, 1, &astr)
	return astr
}

func getAll1(str string, cur int, dep int, balance int,  astr *[]string) {
	if balance < 0 || balance > dep - cur {
		return
	}
	if cur == dep {
		if balance == 0{
			*astr = append(*astr, str)
		}
		return
	} else {
		getAll1(str +"(", cur+1, dep, balance+1,astr)
		getAll1(str +")", cur+1, dep, balance-1,astr)
	}
}


func main()  {
	t := time.Now().UnixMicro()
	generateParenthesis(10)
	//for _, str := range astr{
	//	println(str)
	//}
	println(time.Now().UnixMicro() - t)

	t1 := time.Now().UnixMicro()
	generateParenthesis1(10)
	//for _, str := range astr1{
	//	println(str)
	//}
	println(time.Now().UnixMicro() - t1)
}
