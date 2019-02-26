#define S32_MAX_NUM 0x7fffffff
#define S32_MIN_NUM 0x80000000 * (-1)
typedef struct Node_s{
	int val;
	struct Node_s *pNext;
}Node_t;

void Enqueue(Node_t **Front,Node_t **Rear, int val)
{
	Node_t *NewNode = (Node_t *)malloc(sizeof(Node_t));
	NewNode->val = val;
	NewNode->pNext = NULL;
	if(*Front == NULL && *Rear == NULL) // queue is empty
	{
		*Front = *Rear = NewNode;
	}
	else
	{
		(*Rear)->pNext = NewNode;
		*Rear = NewNode;
	}
}

int Dequeue(Node_t **Front, Node_t **Rear)
{
	if(*Front == NULL)
	{
		return 0;
	}
	Node_t *temp  = *Front;
	int val = (*Front)->val;
	*Front = (*Front)->pNext;
	if(*Front == NULL) // queue is empty update Rear
		*Rear = NULL;
	free(temp);
	return val;
}
int reverse(int x) {

	if(x == 0)
		return 0;
	int Source = x;// > 0 ? x : x * (-1);
	int Result = 0;
	int Remainder = 0;
	Node_t *Front = NULL;
	Node_t *Rear = NULL;
	while(Source != 0)
	{
		Remainder = Source > 0 ? (Source % 10) : -(Source % 10);
		printf("[%s] Remainder = %d \n", __FUNCTION__, Remainder);
		Enqueue(&Front, &Rear, Remainder);
		Source /= 10;
	}
	while(Front != NULL)
	{
		Remainder = Dequeue(&Front, &Rear);
		if(Result > ((S32_MAX_NUM - Remainder) / 10))
		{
			return 0;
		}
		Result = Result * 10 + Remainder;
	}
	Result = x >0 ? Result : Result * (-1);
	return Result;

}
