# Question

Given the head of a singly linked list, reverse the list, and return the reversed list.

Constraints:
- The number of nodes in the list is the range [0, 5000].
- 5000 <= Node.val <= 5000

# Intuition

We want to walk the linked list keeping track of the previous node then utilise a temp variable to keep track of the next node as we swap the pointers around. Standard swap operation.

# Complexity

O(n) time as we just walk the linked list and swap pointers O(1) space as we do it in place so only a few variables being used.
