/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 *
 * https://leetcode.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (38.70%)
 * Total Accepted:    55.1K
 * Total Submissions: 141.5K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * Given two strings s1 and s2, write a function to return true if s2 contains
 * the permutation of s1. In other words, one of the first string's
 * permutations is the substring of the second string.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s1 = "ab" s2 = "eidbaooo"
 * Output: True
 * Explanation: s2 contains one permutation of s1 ("ba").
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:s1= "ab" s2 = "eidboaoo"
 * Output: False
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The input strings only contain lower case letters.
 * The length of both given strings is in range [1, 10,000].
 * 
 * 
 */

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var map1 [26]int
    var map2 [26]int

    for i := range s1 {
        map1[s1[i] - 'a']++
        map2[s2[i] - 'a']++
    }

    for i := 0; i < len(s2) - len(s1); i++ {
        if map1 == map2 {
            return true
        }
        map2[s2[i+len(s1)] -'a']++
        map2[s2[i]-'a']--
    }

    return map1 == map2
}
