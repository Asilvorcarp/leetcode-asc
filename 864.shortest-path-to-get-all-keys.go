/*
 * @lc app=leetcode.cn id=864 lang=golang
 *
 * [864] Shortest Path to Get All Keys
 */

package leetcode

// @lc code=start

import (
	"fmt"
	"strings"
)

// func main() {
// 	t := []string{"@...a", ".###A", "b.BCc"}
// 	fmt.Println(shortestPathAllKeys(t))
// }

// ["@...a", ".###A", "b.BCc"]

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func copyMap(m map[rune]bool) map[rune]bool {
	n := make(map[rune]bool)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func isLock(c byte) bool {
	return c >= 'A' && c <= 'F'
}

func isKey(c byte) bool {
	return c >= 'a' && c <= 'f'
}

func lockToKey(c byte) byte {
	return c - 'A' + 'a'
}

func inString(c rune, s string) bool {
	for _, cc := range s {
		if cc == c {
			return true
		}
	}
	return false
}

func inKeys(c byte, ks int) bool {
	return (1<<(c-'a'))&ks != 0
}

func addToKeySet(ks string, k rune) string {
	for i, c := range ks {
		if k == c {
			return ks
		} else if k < c {
			return ks[:i] + string(k) + ks[i:]
		}
	}
	return ks + string(k)
}

func shortestPathAllKeys(grid []string) int {
	// pState(int) : 0 byte, x byte, y byte, ks byte
	type pState uint32
	row := len(grid)
	col := len(grid[0])
	// board[pState] = dist // this is a 3D board
	board := make([][]byte, row)
	for i := range board {
		board[i] = make([]byte, col)
		// // maybe faster using this instead of visited
		// for j := range board[i] {
		// 	board[i][j] = 255
		// }
	}
	// visited[x][y][ks] = if the pState visited
	visited := make([][][]bool, row)
	for i := range visited {
		visited[i] = make([][]bool, col)
		for j := range visited[i] {
			visited[i][j] = make([]bool, 1<<6)
		}
	}
	// Q: will visited[x*col*64+y*64+ks] be faster? NO?
	queue := make([]pState, 0)

	var fullKeys pState = 0
	for i := range grid {
		for j, c := range grid[i] {
			if c == '@' {
				startPS := pState(i)<<16 | pState(j)<<8 | pState(0)
				// fuck, in C++ I can use inlined function to make it clearer
				board[startPS] = 0
				visited[i][j][0] = true
				queue = append(queue, startPS)
			}
			if c >= 'a' && c <= 'f' {
				fullKeys |= 1 << (c - 'a')
			}
		}
	}

	dir := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		ps := queue[0]
		// fmt.Println("pop", ps)
		queue = queue[1:]
		for _, delta := range dir {
			dis := board[ps]
			// the adjacent pState
			newX := int(ps>>16) + delta[0]
			newY := int(ps>>8|0xff) + delta[1]
			if newX < 0 || newX >= row || newY < 0 || newY >= col {
				continue
			}
			c := grid[newX][newY]

			// fmt.Printf("goto %d %d %c\n", newX, newY, c)

			if c == '#' {
				continue
			}
			newKs := ps | 0xff
			newDis := dis + 1
			if isLock(c) {
				if !inKeys(lockToKey(c), int(ps|0xff)) {
					continue
				}
			} else if isKey(c) {
				newKs |= 1 << (c - 'a')
				if newKs == fullKeys {
					// all keys are collected
					return newDis
				}
			} else if c == '.' || c == '@' {
				// do nothing
			} else {
				fmt.Printf("error - unknown char: %c", c)
				panic("unknown char")
			}
			newPS := pState(newX)<<16 | pState(newY)<<8 | newKs
			if visited[newX][newY][newKs] {
				continue
			}
			board[newPS] = newDis
			queue = append(queue, newPS)
			// fmt.Println("push", pState{newX, newY, newKs})
		}
	}
	// no solution
	return -1
}

// @lc code=end

// my solution v1, TLE, probably right
func shortestPathAllKeys_tooSlow(grid []string) int {
	type pState struct {
		x, y int
		ks   string
	}
	row := len(grid)
	col := len(grid[0])
	// states[x][y][keyStr] = dist
	states := make([][]map[string]int, row)
	for i := range states {
		states[i] = make([]map[string]int, col)
		for j := range states[i] {
			states[i][j] = make(map[string]int)
		}
	}
	queue := make([]pState, 0)

	biggestKey := 'a' - 1
	for i := range grid {
		for j, c := range grid[i] {
			if c == '@' {
				states[i][j][""] = 0
				queue = append(queue, pState{i, j, ""})
			}
			if c >= 'a' && c <= 'f' {
				if c > biggestKey {
					biggestKey = c
				}
			}
		}
	}
	keyNum := int(biggestKey - 'a' + 1)
	// fmt.Println("keyNum", keyNum)

	// directions
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		ps := queue[0]
		// fmt.Println("pop", ps)
		queue = queue[1:]
		for _, delta := range dir {
			dis := states[ps.x][ps.y][ps.ks]
			// the adjacent pState
			newX := ps.x + delta[0]
			newY := ps.y + delta[1]
			if newX < 0 || newX >= row || newY < 0 || newY >= col {
				continue
			}
			c := rune(grid[newX][newY])

			// fmt.Printf("goto %d %d %c\n", newX, newY, c)

			if c == '#' {
				continue
			}
			newKs := ps.ks
			newDis := dis + 1
			if isLock(c) {
				if !inString(lockToKey(c), ps.ks) {
					continue
				}
			} else if isKey(c) {
				newKs = addToKeySet(ps.ks, c)
				if len(newKs) == keyNum {
					// all keys are collected
					return newDis
				}
			} else if c == '.' || c == '@' {
				// do nothing
			} else {
				fmt.Printf("error - unknown char: %c", c)
				panic("unknown char")
			}
			oldDis, exist := states[newX][newY][newKs]
			if exist && newDis > oldDis {
				continue
			}
			states[newX][newY][newKs] = newDis
			queue = append(queue, pState{newX, newY, newKs})
			// fmt.Println("push", pState{newX, newY, newKs})
		}
	}
	// no solution
	return -1
}

// learn cookbook: BFS with pruning
// insane! no store for dists, just a res for queue cleared times counter
func shortestPathAllKeys_cook(grid []string) int {
	if len(grid) == 0 {
		return 0
	}
	board, visited, startx, starty, res, fullKeys := make([][]byte, len(grid)), make([][][]bool, len(grid)), 0, 0, 0, 0
	for i := 0; i < len(grid); i++ {
		board[i] = make([]byte, len(grid[0]))
	}

	for i, g := range grid {
		board[i] = []byte(g)
		for _, v := range g {
			if v == 'a' || v == 'b' || v == 'c' || v == 'd' || v == 'e' || v == 'f' {
				fullKeys |= (1 << uint(v-'a'))
			}
		}
		if strings.Contains(g, "@") {
			startx, starty = i, strings.Index(g, "@")
		}
	}
	for i := 0; i < len(visited); i++ {
		visited[i] = make([][]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			visited[i][j] = make([]bool, 64)
		}
	}
	queue := []int{}
	queue = append(queue, (starty<<16)|(startx<<8))
	visited[startx][starty][0] = true
	for len(queue) != 0 {
		// >> clear the queue, and increase the res (level)
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			state := queue[0]
			queue = queue[1:]
			starty, startx = state>>16, (state>>8)&0xFF
			keys := state & 0xFF
			if keys == fullKeys {
				return res
			}
			for i := 0; i < 4; i++ {
				newState := keys
				nx := startx + dir[i][0]
				ny := starty + dir[i][1]
				if !isInBoard(board, nx, ny) {
					continue
				}
				if board[nx][ny] == '#' {
					continue
				}
				flag, canThroughLock := keys&(1<<(board[nx][ny]-'A')), false
				if flag != 0 {
					canThroughLock = true
				}
				if isLock(board, nx, ny) && !canThroughLock {
					continue
				}
				if isKey(board, nx, ny) {
					newState |= (1 << (board[nx][ny] - 'a'))
				}
				if visited[nx][ny][newState] {
					continue
				}
				queue = append(queue, (ny<<16)|(nx<<8)|newState)
				visited[nx][ny][newState] = true
			}
		}
		res++
	}
	return -1
}
