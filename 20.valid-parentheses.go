/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

package leetcode

// @lc code=start
func isValid(s string) bool {
	a0 := '('
	b0 := ')'
	a1 := '{'
	b1 := '}'
	a2 := '['
	b2 := ']'
	stack := make([]rune, 0)
	for _, c := range s {
		if c == a0 || c == a1 || c == a2 {
			stack = append(stack, c)
		} else if c == b0 || c == b1 || c == b2 {
			if stack[len(stack)] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
