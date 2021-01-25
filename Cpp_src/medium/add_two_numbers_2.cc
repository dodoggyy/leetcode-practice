#include <iostream>
using namespace std;

struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Use general math calculation
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 20 ms, faster than 88.40%
// Memory Usage: 10.4 MB, less than 57.71%
class Solution
{
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
    {
        if (l1 == NULL)
        {
            return l2;
        }
        if (l2 == NULL)
        {
            return l1;
        }
        ListNode *result = new ListNode(0);
        ListNode *head = result;
        int sum = 0;
        while (l1 != NULL || l2 != NULL || sum != 0)
        {
            sum += (l1 == NULL ? 0 : l1->val) + (l2 == NULL ? 0 : l2->val);
            head->next = new ListNode(sum % 10);
            sum /= 10;
            head = head->next;
            if (l1 != NULL)
            {
                l1 = l1->next;
            }
            if (l2 != NULL)
            {
                l2 = l2->next;
            }
        }

        return result->next;
    }
};