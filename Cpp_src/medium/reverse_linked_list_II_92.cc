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
// Runtime: 4 ms, faster than 71.92%
// Memory Usage: 8.2 MB, less than 100.00%
class Solution
{
public:
    ListNode *reverseBetween(ListNode *head, int m, int n)
    {
        ListNode *dummy = new ListNode(-1);
        ListNode *pre = dummy;
        ListNode *cur = NULL;
        pre->next = head;
        for (int i = 0; i < m - 1; i++)
        {
            pre = pre->next;
        }
        cur = pre->next;
        for (int i = m; i < n; i++)
        {
            ListNode *tmp = cur->next;
            cur->next = tmp->next;
            tmp->next = pre->next;
            pre->next = tmp;
        }

        return dummy->next;
    }
};