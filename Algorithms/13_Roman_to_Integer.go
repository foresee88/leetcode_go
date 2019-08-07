//package Algorithms

package main

import "fmt"

/*

13. Roman to Integer
Easy

1317

2728

Favorite

Share
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: "III"
Output: 3
Example 2:

Input: "IV"
Output: 4
Example 3:

Input: "IX"
Output: 9
Example 4:

Input: "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 5:

Input: "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

//num := map[string]int{"I":1, "IV":4, "V":5, "IX":9, "X":10, "XL":40, "L":50, "XC":90, "C":100, "CD":400, "D":500, "CM":900, "M":1000}

func romanToInt(s string) int {

	ret, step, l := 0, false, len(s)
	for i, c := range s {
		if step {
			step = false
			continue
		}
		switch c {

		case 'I':
			if (i < l-1) && s[i+1] == 'V' {
				ret += 4
				step = true
			} else if (i < l-1) && s[i+1] == 'X' {
				ret += 9
				step = true
			} else {
				ret += 1
			}
		case 'V':
			ret += 5
		case 'X':
			if (i < l-1) && s[i+1] == 'L' {
				ret += 40
				step = true
			} else if (i < l-1) && s[i+1] == 'C' {
				ret += 90
				step = true
			} else {
				ret += 10
			}
		case 'L':
			ret += 50
		case 'C':
			if (i < l-1) && s[i+1] == 'D' {
				ret += 400
				step = true
			} else if (i < l-1) && s[i+1] == 'M' {
				ret += 900
				step = true
			} else {
				ret += 100
			}

		case 'D':
			ret += 500
		case 'M':
			ret += 1000
		}
	}

	return ret
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
