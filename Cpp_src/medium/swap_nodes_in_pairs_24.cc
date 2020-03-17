#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 72.88%
// Memory Usage: 8 MB, less than 100.00%
class Solution
{
public:
    ListNode *swapPairs(ListNode *head)
    {
        ListNode *dummy = new ListNode(-1);
        dummy->next = head;
        ListNode *cur = dummy;

        while (cur->next && cur->next->next)
        {
            ListNode *l1 = cur->next, *l2 = cur->next->next;
            ListNode *tmp = l2->next;
            l1->next = tmp;
            l2->next = l1;
            cur->next = l2;

            cur = l1;
        }

        return dummy->next;
    }
};

// Use recursive:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 72.88%
// Memory Usage: 8 MB, less than 100.00%
class Solution2
{
public:
    ListNode *swapPairs(ListNode *head)
    {
        if (head == NULL || head->next == NULL)
        {
            return head;
        }

        ListNode *tmp = head->next;
        head->next = swapPairs(head->next->next);
        tmp->next = head;

        return tmp;
    }
};