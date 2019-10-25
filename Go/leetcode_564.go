/*
 * @lc app=leetcode id=564 lang=golang
 *
 * [564] Find the Closest Palindrome
 *
 * https://leetcode.com/problems/find-the-closest-palindrome/description/
 *
 * algorithms
 * Hard (19.13%)
 * Total Accepted:    16.3K
 * Total Submissions: 85.1K
 * Testcase Example:  '"1"'
 *
 * Given an integer n, find the closest integer (not including itself), which
 * is a palindrome. 
 * 
 * The 'closest' is defined as absolute difference minimized between two
 * integers.
 * 
 * Example 1:
 * 
 * Input: "123"
 * Output: "121"
 * 
 * 
 * 
 * Note:
 * 
 * The input n is a positive integer represented by string, whose length will
 * not exceed 18.
 * If there is a tie, return the smaller one as answer.
 * 
 * 
 */

 
import (
	"fmt"
	"strings"
	"strconv"
)

func isPalindrome(s string) bool {
    j := len(s) - 1
    for i := 0; i < j; i++ {
        if s[i] != s[j] {
            return false
        }
        j--
    }
    return true
}

func genMirrorPalindrome(n string) string {
	var buf strings.Builder
	mid := len(n) / 2
	left := n[:mid]
	buf.WriteString(left)
	if len(n) % 2 == 1 {
		buf.WriteByte(n[mid])
	}
	for i := len(left)-1; i >= 0; i-- {
		buf.WriteByte(left[i])
	}
	return buf.String()
}

func bruteGen(n int, offset int, stopAt int) string {
	fmt.Printf("og: %d, d: %d, stop: %d\n", n, offset, stopAt)
	for i := n+offset; i != stopAt; i=i+offset {
		str := strconv.Itoa(i)
		if isPalindrome(str) {
			return str
		}
	}
	return ""
}

func getDelta(n int, m string) int {
	m2, _ := strconv.Atoi(m)
	return m2 - n
}
func nearestPalindromic(n string) string{
	nVal, _ := strconv.Atoi(n)
	
	if len(n) == 1 {
		if nVal == 0 {
			return "1"
		} else {
			return strconv.Itoa(nVal-1)
		}
	}
	
	var loStr, hiStr string
	var lo, hi int
	
	ez := genMirrorPalindrome(n)
	d := getDelta(nVal, ez)
	fmt.Printf("EZ: %s, d: %d\n", ez, d)
	if d < 0 {
		loStr = ez
		lo = -d
	} else if 0 < d {
		hiStr = ez
		hi = d
	}
	
	if loStr == "" {
		var stopAt int
		if hiStr != "" {
			stopAt = nVal-hi-1
		} else {
			stopAt = -1
		}
		loStr = bruteGen(nVal, -1, stopAt)
		if loStr != "" {
			tmp, _ := strconv.Atoi(loStr)
			lo = nVal - tmp
		}
	}
	fmt.Printf("LoStr: %s, Lo: %d\n", loStr, lo)
	
	if hiStr == "" {
		hiStr = bruteGen(nVal, 1, nVal+lo)
		if hiStr != "" {
			tmp, _ := strconv.Atoi(hiStr)
			hi = tmp - nVal
		}
	}
	fmt.Printf("HiStr: %s, Hi: %d\n", hiStr, hi)
	
    if lo == 0 || (0 < hi && hi < lo) {
        return hiStr
    } else {
        return loStr
    }
}
