#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Use pointers
// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 16 ms, faster than 70.94%
// Memory Usage: 9.7 MB, less than 71.43%
class Solution
{
public:
    ListNode *oddEvenList(ListNode *head)
    {
        if (head == NULL)
        {
            return head;
        }
        ListNode *odd = head, *even = head->next, *even_head = even;
        while (even != NULL && even->next != NULL)
        {
            odd->next = odd->next->next;
            even->next = even->next->next;
            odd = odd->next;
            even = even->next;
        }
        odd->next = even_head;

        return head;
    }
};