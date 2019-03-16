/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
/**
 * Return an array of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */

 typedef struct Node_Tree_s
 {
 	struct TreeNode *Node;
 	struct Linklist *pNext;
 }Node_Tree_t;

typedef struct Node_Int_s
{
	int val;
	struct Node_Int *pNext;
}Node_Int_t;

void Push_stack(Node_Tree_t **StackTop, struct TreeNode *Node)
{
	Node_Tree_t *temp = (Node_Tree_t *)malloc(sizeof(Node_Tree_t));
	if(!(*StackTop)) // Stack top is null
	{
		temp->pNext = NULL;
	}
	else
	{
		temp->pNext = (*StackTop);
	}

	temp->Node = Node;
	*StackTop = temp; // modify stack_top stroed address

}

struct TreeNode *Pop_stack(Node_Tree_t **StackTop)
{
	Node_Tree_t *temp = (*StackTop);
	struct TreeNode *Node = NULL;
	*StackTop = (*StackTop)->pNext;
	Node = temp->Node;
	free(temp);
	return Node;
}
void Enqueue(Node_Int_t **Front, int val)
{
	Node_Int_t *temp = (Node_Int_t *)malloc(sizeof(Node_Int_t));
	temp->pNext = NULL;
	temp->val = val;
	if((*Front)) // Queue is not empty
	{
		(*Front)->pNext = temp;
	}
	*Front = temp;
}

int Dequeue(Node_Int_t **Rear)
{
	Node_Int_t *temp = *Rear;
	int val = (*Rear)->val;
	*Rear = (*Rear)->pNext;
	free(temp);
	return val;
}

//1. use stack to store traversal node.
//2. use queue to store acquire node's val order.
int* preorderTraversal(struct TreeNode* root, int* returnSize) {
    
    if(!root)
    {
    	*returnSize = 0;
    	return NULL;
    }

    Node_Tree_t *StackTop = NULL;
    Node_Int_t *QueueFront = NULL;
    Node_Int_t *QueueRear = NULL;
    Push_stack(&StackTop, root);

    while(StackTop)
    {
    	struct TreeNode *Node = Pop_stack(&StackTop);
    	if(QueueFront == NULL)
    	{
    		Enqueue(&QueueFront, Node->val);
    		QueueRear = QueueFront;
    	}
    	else
    		Enqueue(&QueueFront, Node->val);

    	//Push TreeNode to stack by traversal order
    	if(Node->right)
    		Push_stack(&StackTop, Node->right);
    	if(Node->left)
    		Push_stack(&StackTop, Node->left);

    	*returnSize = *returnSize + 1;
    }

    int *Result = (int *)malloc(sizeof(int) * (*returnSize));
    unsigned int idx = 0;
    while(QueueRear)
    {
    	if(QueueRear == QueueFront)
    	{
    		Result[idx] = Dequeue(&QueueRear);
    		QueueFront = QueueRear;	
    	}
    	else
    		Result[idx] = Dequeue(&QueueRear);

    	idx++;
    }
    return Result;
}