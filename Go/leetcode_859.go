/*
 * @lc app=leetcode id=859 lang=golang
 *
 * [859] Buddy Strings
 *
 * https://leetcode.com/problems/buddy-strings/description/
 *
 * algorithms
 * Easy (27.83%)
 * Total Accepted:    32K
 * Total Submissions: 115K
 * Testcase Example:  '"ab"\n"ba"'
 *
 * Given two strings A and B of lowercase letters, return true if and only if
 * we can swap two letters in A so that the result equals B.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * 
 * Input: A = "ab", B = "ba"
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: A = "ab", B = "ab"
 * Output: false
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: A = "aa", B = "aa"
 * Output: true
 * 
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: A = "aaaaaaabc", B = "aaaaaaacb"
 * Output: true
 * 
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: A = "", B = "aa"
 * Output: false
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * 0 <= A.length <= 20000
 * 0 <= B.length <= 20000
 * A and B consist only of lowercase letters.
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 */
func buddyStrings(A string, B string) bool {
    if len(A) != len(B) {
        return false
    }
    
    if A == B {
        seen := map[rune]bool{}
        for _, c := range A {
            _, ok := seen[c]
            if ok {
                return true
            }
            seen[c] = true
        }
        return false
    }
    
    pairs := [][]byte{}
    for i := 0; i < len(A); i++ {
        if A[i] != B[i] {
            pairs = append(pairs, []byte{A[i], B[i]})
        }
    }
    
    return len(pairs) == 2 && pairs[0][0] == pairs[1][1] && pairs[0][1] == pairs[1][0]
}
