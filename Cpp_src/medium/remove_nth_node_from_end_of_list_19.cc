#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Use Two pass algorithm
// Time Complexity: O(L)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 86.13%
// Memory Usage: 8.5 MB, less than 76.32%
class Solution
{
public:
    ListNode *removeNthFromEnd(ListNode *head, int n)
    {
        if (head->next == NULL)
        {
            return NULL;
        }
        ListNode *cur = head;
        int length = 0;

        while (cur)
        {
            length++;
            cur = cur->next;
        }
        if (n == length)
        {
            return head->next;
        }

        cur = head;
        for (int i = 0; i < (length - n - 1); i++)
        {
            cur = cur->next;
        }
        ListNode *tmp = cur->next;
        cur->next = cur->next->next;
        delete tmp;

        return head;
    }
};

// Use One pass algorithm
// Time Complexity: O(L)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 86.13%
// Memory Usage: 8.5 MB, less than 80.26%
class Solution2
{
public:
    ListNode *removeNthFromEnd(ListNode *head, int n)
    {
        if (head->next == NULL)
        {
            return NULL;
        }
        ListNode *fast = head;
        ListNode *slow = head;

        while (n--)
        {
            fast = fast->next;
        }
        if (fast == NULL)
        {
            return head->next;
        }
        while (fast->next)
        {
            slow = slow->next;
            fast = fast->next;
        }
        ListNode *tmp = slow->next;
        slow->next = slow->next->next;
        delete tmp;

        return head;
    }
};