/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (32.76%)
 * Total Accepted:    428.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 * 
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 * 
 * Example 1:
 * 
 * 
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "race a car"
 * Output: false
 * 
 * 
 */

import (
    "regexp"
    "strings"
)

func sanatizeString(s string) string {
    reg, _ := regexp.Compile("[^a-zA-Z0-9]")
    s = reg.ReplaceAllString(s, "")
    return strings.ToLower(s)
}

func isPalindrome(s string) bool {
    s = sanatizeString(s)
    if s == "" {
        return true
    }
    j := len(s) - 1
    for i:= 0; i < j; i++ {
        if s[i] != s[j] {
            return false
        }
        j--
    }

    return true
}
