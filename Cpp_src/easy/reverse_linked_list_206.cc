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
// Runtime: 8 ms, faster than 77.57%
// Memory Usage: 9.1 MB, less than 100.00%
class Solution
{
public:
    ListNode *reverseList(ListNode *head)
    {
        ListNode *pre = NULL;
        ListNode *cur = head;

        while (cur != NULL)
        {
            ListNode *tmp = cur->next;
            cur->next = pre;
            pre = cur;
            cur = tmp;
        }
        return pre;
    }
};

// Use recursive
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 77.57%
// Memory Usage: 9.2 MB, less than 82.44%
class Solution2
{
public:
    ListNode *reverseList(ListNode *head)
    {
        if (head == NULL || head->next == NULL)
        {
            return head;
        }
        ListNode *tmp = reverseList(head->next);
        head->next->next = head;
        head->next = NULL;
        return tmp;
    }
};