/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (35.95%)
 * Total Accepted:    368.5K
 * Total Submissions: 1M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 * 
 * Example 1:
 * 
 * 
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 * 
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 * 
 */

import (
    "fmt"
    "sort"
)

type intervalSlice [][]int

func (s intervalSlice) Len() int {
    return len(s)
}

func (s intervalSlice) Swap (i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s intervalSlice) Less(i, j int) bool {
    return s[i][0] < s[j][0]
}

func max (x,y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}

func merge(intervals intervalSlice) intervalSlice {
    if len(intervals) == 0 {
        return intervals
    }

    newIntervals := intervalSlice{}

    sort.Sort(intervals)
    lastInterval := intervals[0]
    for i := 1; i < len(intervals); i++ {
        currInterval := intervals[i]
        if lastInterval[1] >= currInterval[0] {
            lastInterval[1] = max(lastInterval[1], currInterval[1])
        } else {
            newIntervals = append(newIntervals, lastInterval)
            lastInterval = currInterval
        }
    }

    return append(newIntervals, lastInterval)
}
