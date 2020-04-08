# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        p = head
        cnt = 0
        while p != None:
            p = p.next
            cnt += 1

        p = head
        for _ in range(cnt//2):
            p = p.next

        return p


def build_linked_list(array):
    if len(array) == 0:
        return None

    root = prev = ListNode(array[0])
    for el in array[1:]:
        node = ListNode(el)
        prev.next = node
        prev = node

    return root


if __name__ == "__main__":
    solver = Solution()

    root = build_linked_list([1,2,3,4,5])
    middle = solver.middleNode(root)
    assert middle.val == 3

    root = build_linked_list([1,2,3,4,5,6])
    middle = solver.middleNode(root)
    assert middle.val == 4

    root = build_linked_list([2])
    middle = solver.middleNode(root)
    assert middle.val == 2
