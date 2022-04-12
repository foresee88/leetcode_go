package main

import (
	"fmt"
	"strings"
)

/*
30. Substring with Concatenation of All Words
Hard

1875

1825

Add to List

Share
You are given a string s and an array of strings words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.

You can return the answer in any order.



Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]


Constraints:

1 <= s.length <= 104
s consists of lower-case English letters.
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] consists of lower-case English letters.
Accepted
249,801
Submissions
892,217
*/

// 获取排列的算法参考 https://www.cnblogs.com/autosar/archive/2012/04/08/2437799.html

func findSubstring(s string, words []string) []int {
	var as []string
	var ai []int
	perm(words, 0, len(words)-1, &as)



	//m := make(map[string]string)
	//for _, sub := range as {
	//	m[sub] = sub
	//}
	//if len(m) != len(as) {
	//	as = []string{}
	//	for _, v := range m {
	//		as = append(as, v)
	//	}
	//}
	fmt.Println(as)
	fmt.Println(s)

	//for _, sub := range as {
	//	m, ss := 0, s
	//	for len(ss) >= len(sub) {
	//		i := strings.Index(ss, sub)
	//		if i == -1{
	//			break
	//		}
	//		ai = append(ai, m+i)
	//		ss = ss[i+1:]
	//		m += i+1
	//	}
	//}
	return ai
}

func perm(words []string, s int, e int, ret *[]string) {
	if s > e {
		*ret = append(*ret, strings.Join(words, ""))
	} else {
		for i := s; i <= e; i++ {
			swap(words, s, i)
			perm(words, s+1, e, ret)
			swap(words, s, i)
		}
	}
}

func swap(o []string, i int, j int) {
	tmp := o[i]
	o[i] = o[j]
	o[j] = tmp
}

func main() {
	//ai := findSubstring("foobarfoobar", []string{"foo", "bar"})
	//fmt.Println(ai)
	//ai1 := findSubstring("a", []string{"a"})
	//fmt.Println(ai1)
	ai2 := findSubstring("pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbcjjivtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbjjkgcztkcbwagtcnrtqryuqixtzhaakjlurnumzyovawrcjiwabuwretmdamfkxrgqgcdgbrdbnugzecbgyxxdqmisaqcyjkqrntxqmdrczxbebemcblftxplafnyoxqimkhcykwamvdsxjezkpgdpvopddptdfbprjustquhlazkjfluxrzopqdstulybnqvyknrchbphcarknnhhovweaqawdyxsqsqahkepluypwrzjegqtdoxfgzdkydeoxvrfhxusrujnmjzqrrlxglcmkiykldbiasnhrjbjekystzilrwkzhontwmehrfsrzfaqrbbxncphbzuuxeteshyrveamjsfiaharkcqxefghgceeixkdgkuboupxnwhnfigpkwnqdvzlydpidcljmflbccarbiegsmweklwngvygbqpescpeichmfidgsjmkvkofvkuehsmkkbocgejoiqcnafvuokelwuqsgkyoekaroptuvekfvmtxtqshcwsztkrzwrpabqrrhnlerxjojemcxel", []string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty"})
	fmt.Println(ai2)
}
