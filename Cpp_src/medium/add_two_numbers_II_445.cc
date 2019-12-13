#include <iostream>
#include <stack>
using namespace std;

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Use stack
// Time Complexity: O(MAX(m,n))
// Space Complexity: O(MAX(m,n))
// Runtime: 24 ms, faster than 75.67%
// Memory Usage: 13 MB, less than 51.85%
class Solution
{
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
    {
        ListNode *result = new ListNode(0);
        ListNode *head = result;
        stack<int> stack1;
        stack<int> stack2;

        while (l1)
        {
            stack1.push(l1->val);
            l1 = l1->next;
        }

        while (l2)
        {
            stack2.push(l2->val);
            l2 = l2->next;
        }

        int carry = 0;
        while (!stack1.empty() || !stack2.empty() || carry != 0)
        {
            int num1 = 0;
            int num2 = 0;
            if (!stack1.empty())
            {
                num1 = stack1.top();
                stack1.pop();
            }

            if (!stack2.empty())
            {
                num2 = stack2.top();
                stack2.pop();
            }
            int sum = num1 + num2 + carry;
            ListNode *tmp = new ListNode(sum % 10);
            tmp->next = head->next;
            head->next = tmp;

            carry = sum / 10;
        }

        return result->next;
    }
};