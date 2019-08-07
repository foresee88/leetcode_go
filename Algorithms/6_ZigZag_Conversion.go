//package ___ZigZag_Conversion
package main

import "fmt"

/*

6. ZigZag Conversion
Medium

1096

3379

Favorite

Share
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

*/

//思路：构建一个二维数组，按照规则把符合条件的元素放到对应的位置上
func convert(s string, numRows int) string {
	if s == "" || numRows <= 1 {
		return s
	}

	// 计算二维数组所需要的列数，行数 + 对角线上的元素个数（等于行数-2) 形成一个列计算单位
	maxCol := ((len(s) / (numRows + numRows - 2)) + 1) * numRows
	str := make([][]string, numRows)
	for i := range str {
		str[i] = make([]string, maxCol)
	}

	ret := ""
	count := 0
	for i := 0; i < maxCol; i++ {
		for j := 0; j < numRows; j++ {
			if i%(numRows-1) == 0 || (i+j)%(numRows-1) == 0 {
				str[j][i] = string(s[count])
				count++
				if count >= len(s) {
					break
				}
			}
		}
		if count >= len(s) {
			break
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < maxCol; j++ {
			if str[i][j] != "" {
				ret += str[i][j]
			}
		}
	}

	return ret
}

func main(){
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
