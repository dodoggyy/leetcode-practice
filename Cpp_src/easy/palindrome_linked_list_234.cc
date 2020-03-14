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
// Runtime: 20 ms, faster than 93.63%
// Memory Usage: 11.9 MB, less than 100.00%
class Solution
{
public:
    bool isPalindrome(ListNode *head)
    {
        if (head == NULL)
        {
            return true;
        }
        ListNode *slow = head;
        ListNode *fast = head;

        while (fast && fast->next)
        {
            slow = slow->next;
            fast = fast->next->next;
        }
        slow = reverse(slow);

        while (slow)
        {
            if (slow->val != head->val)
            {
                return false;
            }
            slow = slow->next;
            head = head->next;
        }

        return true;
    }

    ListNode *reverse(ListNode *head)
    {
        ListNode *pre = NULL;
        ListNode *cur = head;

        while (cur)
        {
            ListNode *tmp = cur->next;
            cur->next = pre;
            pre = cur;
            cur = tmp;
        }

        return pre;
    }
};