#include <iostream>
#include <vector>
using namespace std;

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Split Input List:
// Time Complexity: O(N+k)
// Space Complexity:O(k)
// Runtime: 12 ms, faster than 74.19%
// Memory Usage: 7.7 MB, less than 100.00%
class Solution
{
public:
    vector<ListNode *> splitListToParts(ListNode *root, int k)
    {
        vector<ListNode *> result(k, NULL);
        int size = 0;
        ListNode *cur = root;
        while (cur)
        {
            size++;
            cur = cur->next;
        }

        int len = size / k, rem = size % k;

        cur = root;
        for (int i = 0; i < k; i++, rem--)
        {
            if (cur == NULL)
            {
                break;
            }
            ListNode *pre = NULL;
            result[i] = cur;
            for (int j = 0; j < len + (rem > 0 ? 1 : 0); j++)
            {
                pre = cur;
                cur = cur->next;
            }
            pre->next = NULL;
        }

        return result;
    }
};