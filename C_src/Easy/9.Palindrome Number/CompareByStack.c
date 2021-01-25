typedef struct stNode
{
	int val;
	struct stNode *pNext;
}Node;

void StackPush(Node **pTop, int val)
{
	Node *pTmp = (Node *)malloc(sizeof(Node));
	pTmp->pNext = *pTop;
	pTmp->val = val;
	*pTop = pTmp;
	return;
}

int StackPop(Node **pTop)
{
	int Rel = (*pTop)->val;
	Node *Tmp = (*pTop);
	*pTop = (*pTop)->pNext;
	free(Tmp);
	return Rel;
}

bool isPalindrome(int x){

	if(!(x) ||  (x < 0))
	{
		return 0;
	}

	int Tmp = x;
	int digital = 0;
	unsigned char bIsPalindrome = 1;
	Node *pTop = NULL;
	while(Tmp)
	{
		digital = Tmp % 10;
		StackPush(&pTop, digital);
		Tmp /= 10 ;
	}

	Tmp = x;
	while(Tmp)
	{
		digital = Tmp % 10;
		if(digital != StackPop(&pTop))
		{
			bIsPalindrome = 0;
		}
		Tmp /= 10;
	}
	return bIsPalindrome;
}
