# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __str__(self):
        return f"{self.val}{' -> ' + str(self.next) if self.next else ''}"

class Solution:
    def reverse_list(self, l):
        cur = None
        n1 = l
        n2 = l.next
        while n2:
            n1.next = cur
            cur = n1
            n1 = n2
            n2 = n2.next
        n1.next = cur
        return n1
        
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        # l1 = self.reverse_list(l1)
        # l2 = self.reverse_list(l2)
        cur_l1 = l1
        cur_l2 = l2
        cur_l3 = None
        head_l3 = None
        carry = 0
        while True:
            print(cur_l3)
            if cur_l1 == None and cur_l2 == None:
                break
            cur_l3 = ListNode(val=carry, next=head_l3)
            if  cur_l1:
                cur_l3.val += cur_l1.val
                cur_l1 = cur_l1.next
            if  cur_l2:
                cur_l3.val += cur_l2.val
                cur_l2 = cur_l2.next
            carry = int(cur_l3.val / 10)
            cur_l3.val %= 10
            head_l3 = cur_l3
        if carry:
            cur_l3 = ListNode(val=carry, next=cur_l3)
        return self.reverse_list(cur_l3)

[2,4,9]
[5,6,4,9]

s = Solution()
l1  = ListNode(2, ListNode(4, ListNode(9)))
l2  = ListNode(5, ListNode(6, ListNode(4, ListNode(9))))
print(l1)
print(l2)
print(s.addTwoNumbers(l1, l2))