#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Use iterative
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 74.28%
// Memory Usage: 9.4 MB, less than 39.62%
class Solution
{
public:
    ListNode *deleteDuplicates(ListNode *head)
    {
        if (head == NULL || head->next == NULL)
        {
            return head;
        }
        ListNode *cur = head;

        while (cur->next)
        {
            if (cur->val == cur->next->val)
            {
                ListNode *tmp = cur->next;
                cur->next = cur->next->next;
                delete tmp;
            }
            else
            {
                cur = cur->next;
            }
        }

        return head;
    }
};

// Use recursive
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 74.28%
// Memory Usage: 9.1 MB, less than 100.00%
class Solution2
{
public:
    ListNode *deleteDuplicates(ListNode *head)
    {
        if (head == NULL || head->next == NULL)
        {
            return head;
        }
        head->next = deleteDuplicates(head->next);
        return (head->val == head->next->val) ? head->next : head;
    }
};