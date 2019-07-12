/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (31.35%)
 * Total Accepted:    925K
 * Total Submissions: 3M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 * 
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 * 
 * Example:
 * 
 * 
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

func addNodes(x, y *ListNode, c int) (sum, carry int) {
    sum = x.Val + y.Val + c
    carry = sum / 10
    sum %= 10
    return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var head, res *ListNode
    zeroNode := &ListNode{
        Val: 0,
    }
    
    for l1 != nil || l2 != nil {
        if res == nil {
            res = &ListNode{}
            head = res
        } else {
            res.Next = &ListNode{}
            res = res.Next
        }

        if l1 == nil {
            l1 = zeroNode
        }
        if l2 == nil {
            l2 = zeroNode
        }

        res.Val, carry = addNodes(l1, l2, carry)
        l1 = l1.Next
        l2 = l2.Next
    }

    if carry > 0 {
        res.Next = &ListNode{
            Val: carry,
        }
    }
    return head
}
