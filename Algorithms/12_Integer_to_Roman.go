//package _2__Integer_to_Roman
package main

import "fmt"

/*
12. Integer to Roman
Medium

617

1942

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
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: 3
Output: "III"
Example 2:

Input: 4
Output: "IV"
Example 3:

Input: 9
Output: "IX"
Example 4:

Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 5:

Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

*/

/*
func addCh(ss []string, s string, n int) []string {
	for i := 0; i < n; i++ {
		ss = append(ss, s)
	}
	return ss
}

func intToRoman(num int) string {
	s := make([]string, 0)
	nM := num / 1000
	s = addCh(s, "M", nM)

	num = num % 1000
	nC := num / 100
	switch nC {
	case 4:
		s = addCh(s, "CD", 1)
		break
	case 5:
		s = addCh(s, "D", 1)
		break
	case 9:
		s = addCh(s, "CM", 1)
		break
	default:
		if nC < 4 {
			s = addCh(s, "C", nC)
		} else {
			s = addCh(s, "D", 1)
			s = addCh(s, "C", nC-5)
		}

	}

	num = num % 100
	nX := num / 10
	switch nX {
	case 4:
		s = addCh(s, "XL", 1)
		break
	case 5:
		s = addCh(s, "L", 1)
		break
	case 9:
		s = addCh(s, "XC", 1)
		break
	default:
		if nX < 4 {
			s = addCh(s, "X", nX)
		} else {
			s = addCh(s, "L", 1)
			s = addCh(s, "X", nX-5)
		}

	}

	num = num % 10
	nI := num
	switch nI {
	case 4:
		s = addCh(s, "IV", 1)
		break
	case 5:
		s = addCh(s, "V", 1)
		break
	case 9:
		s = addCh(s, "IX", 1)
		break
	default:
		if nI < 4 {
			s = addCh(s, "I", nI)
		} else {
			s = addCh(s, "V", 1)
			s = addCh(s, "I", nI-5)
		}
	}

	var str string
	for _,v := range s {
		str += v
	}
	return str
}
*/


/*
func addCh(bs []byte, b byte, n int) []byte {
	for i := 0; i < n; i++ {
		bs = append(bs, b)
	}
	return bs
}

func convert(bs []byte, num int, digit int, a byte, b byte, c byte) []byte{
	n := num/digit
	switch n {
	case 4:
		bs = append(bs, a, b)
		break
	case 5:
		bs = append(bs, b)
		break
	case 9:
		bs = append(bs, a, c)
		break
	default:
		if n < 4 {
			bs = addCh(bs, a, n)
		} else {
			bs = append(bs, b)
			bs = addCh(bs, a, n-5)
		}
	}
	return bs
}

func intToRoman(num int) string {
	bs := make([]byte, 0)
	nM := num / 1000
	bs = addCh(bs, 'M', nM)

	num = num % 1000
	bs = convert(bs, num, 100, 'C', 'D', 'M')

	num = num % 100
	bs = convert(bs, num, 10,'X', 'L', 'C')

	num = num % 10
	bs = convert(bs, num, 1,'I', 'V', 'X')

	s := string(bs)
	return s
}
*/


func addCh(bsp *[]byte, b byte, n int) {
	for i := 0; i < n; i++ {
		*bsp = append(*bsp, b)
	}
}

func convert12(bsp *[]byte, num int, digit int, a byte, b byte, c byte) {
	n := num/digit
	switch n {
	case 4:
		*bsp = append(*bsp, a, b)
		break
	case 5:
		*bsp = append(*bsp, b)
		break
	case 9:
		*bsp = append(*bsp, a, c)
		break
	default:
		if n < 4 {
			addCh(bsp, a, n)
		} else {
			*bsp = append(*bsp, b)
			addCh(bsp, a, n-5)
		}
	}
}

func intToRoman(num int) string {
	bs := make([]byte, 0)
	nM := num / 1000
	addCh(&bs, 'M', nM)

	num = num % 1000
	convert12(&bs, num, 100, 'C', 'D', 'M')

	num = num % 100
	convert12(&bs, num, 10,'X', 'L', 'C')

	num = num % 10
	convert12(&bs, num, 1,'I', 'V', 'X')

	s := string(bs)
	return s
}



func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
