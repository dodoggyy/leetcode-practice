/*
Result:
Runtime: 4 ms, faster than 62.91% of C online submissions for Swap Nodes in Pairs.
Memory Usage: 7.2 MB, less than 50.00% of C online submissions for Swap Nodes in Pairs.
*/

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

/*
	For this case we need 3 ptr
	1. pointer to previous node
	2. pointer to current node
	3. pointer to next node
*/

struct ListNode* swapPairs(struct ListNode* head){

	if((!head))
		return head;


	struct ListNode* pPrev = NULL;
	struct ListNode* pCur = head;
	struct ListNode* pNext = head->next;

	while(((pCur != NULL) && (pNext != NULL)))
	{
		if(pPrev != NULL)
		{
			pPrev->next = pNext;
			pCur->next = pNext->next;
			pNext->next = pCur;
		}
		else // first node swap case, so need move head ptr to next node
		{
            head = pNext;
			pCur->next = pNext->next;
			pNext->next = pCur;
		}

		pPrev = pCur;
		pCur = pCur->next;
		pNext = (pCur != NULL) ? pCur->next : NULL;
	}

	return head;
}

