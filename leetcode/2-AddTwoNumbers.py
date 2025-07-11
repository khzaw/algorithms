from typing import Optional

class ListNode:
    def __init__(self, x:int) -> None:
        self.val = x
        self.next: Optional[ListNode] = None


class Solution:
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """

        result_root = n = ListNode(0)
        carry = 0

        while l1 or l2 or carry:
            v1 = v2 = 0

            if l1:
                v1 = l1.val
                l1 = l1.next

            if l2:
                v2 = l2.val
                l2 = l2.next

            carry, val = divmod(v1 + v2 + carry, 10)

            n.next = ListNode(val)
            n = n.next

        return result_root.next
