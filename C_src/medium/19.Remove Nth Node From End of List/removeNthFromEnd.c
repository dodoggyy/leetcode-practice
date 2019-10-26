/*
Runtime: 0 ms, faster than 100.00% of C online submissions for Remove Nth Node From End of List.
Memory Usage: 7 MB, less than 100.00% of C online submissions for Remove Nth Node From End of List.
*/


/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

/*

	Assume total length of linklist is k
	If we want to remove the n-th node from the end of list.
	The target node actual exist in (k-n)-th node.

	1. counting total lenght of list.
	2. calcaulte target node number.
	3. remove target node.
*/

struct ListNode* removeNthFromEnd(struct ListNode* head, int n){

	if(!head)
	{
		return NULL;
	}
	int Lenght = 0;
	int TargetIdx = 0;
	int Idx = 0;
	struct ListNode* Tmp = head;
	struct ListNode* TargetNode = NULL;
	while(Tmp)
	{
		Lenght++;
		Tmp = Tmp->next;
	}

	TargetIdx = Lenght - n;
	if(TargetIdx == 0) // Target node in head
	{
		TargetNode = head;
		head = head->next;
	}
	else
	{
        Tmp = head;
		while(Tmp)
		{
			if(Idx == (TargetIdx -1)) // locate previous node of target node
			{
				break;
			}
			Tmp = Tmp->next;
			Idx++;
		}
		// Remove target node from list
		TargetNode = Tmp->next;
		Tmp->next = TargetNode->next;
	}
	free(TargetNode);
	return head;
}
