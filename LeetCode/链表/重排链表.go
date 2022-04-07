package main

/*给定一个单链表 L 的头节点 head ，单链表 L 表示为：
L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

// 思路：先通过快慢指针找到中间节点，然后从中间节点断开，反转后面的链表，然后合并

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}

}