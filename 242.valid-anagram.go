/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] Valid Anagram
 */

/**
 * hint: 打表
 */
package leetcode

// @lc code=start
func isAnagram(s string, t string) bool {
	table := make([]int, 26)
	var a rune = 'a'
	for _, c := range s {
		table[c-a]++
	}
	for _, c := range t {
		table[c-a]--
	}
	for _, n := range table {
		if n != 0 {
			return false
		}
	}
	return true
}

// learn:

func isAnagram1(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

// map
func isAnagram2(s string, t string) bool {
	hash := map[rune]int{}
	for _, value := range s {
		hash[value]++
	}
	for _, value := range t {
		hash[value]--
	}
	for _, value := range hash {
		if value != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
