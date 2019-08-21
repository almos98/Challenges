/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (52.41%)
 * Total Accepted:    374.1K
 * Total Submissions: 706.3K
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an array of integers, find if the array contains any duplicates.
 * 
 * Your function should return true if any value appears at least twice in the
 * array, and it should return false if every element is distinct.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3,1]
 * Output: true
 * 
 * Example 2:
 * 
 * 
 * Input: [1,2,3,4]
 * Output: false
 * 
 * Example 3:
 * 
 * 
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 * 
 */

import "sort"

type numSlice []int

func (s numSlice) Len() int {
    return len(s)
}

func (s numSlice) Swap (i,j int) {
    s[i], s[j] = s[j], s[i]
}

func (s numSlice) Less (i, j int) bool {
    return s[i] < s[j]
}

func containsDuplicate(nums numSlice) bool {
    sort.Sort(nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}
