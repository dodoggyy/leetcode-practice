/* Result:
Runtime: 8 ms, faster than 86.93% of C online submissions for Find First and Last Position of Element in Sorted Array.
Memory Usage: 8.7 MB, less than 50.00% of C online submissions for Find First and Last Position of Element in Sorted Array.
*/

/*
Search target index by binary tree.
3 of case need handle
1. root < target, then traversal left and right subtree if subtree exist.
2. root = target, then update first(min) or last(max) Idx and traversal left and right subtree if subtree exist.
3. root > target stop traversal subtree.
*/

static void TraversalBT(int *nums, int numsSize, int CurIdx, int target, int *pMin, int *pMax)
{
	if(nums[CurIdx] <= target)
	{
		if(nums[CurIdx] == target)
		{
			*pMin = ((*pMin >= CurIdx) || (*pMin == -1)) ? CurIdx : *pMin;
			*pMax = ((*pMax <= CurIdx) || (*pMax == -1)) ? CurIdx : *pMax;
		}
		//Traversal left subtree
		if(numsSize > ((CurIdx * 2) + 1))
			TraversalBT(nums, numsSize, ((CurIdx * 2) + 1), target, pMin, pMax);
		//Traversal right subtree
		if(numsSize > ((CurIdx * 2) + 2))  
			TraversalBT(nums, numsSize, ((CurIdx * 2) + 2), target, pMin, pMax);
	}
	return;
}

int* searchRange(int* nums, int numsSize, int target, int* returnSize){

	*returnSize = 2;
	int *pOutputIdx = (int *)malloc((*returnSize) * sizeof(int));
	memset(pOutputIdx, -1, ((*returnSize) * sizeof(int)));
	if((!nums) || (numsSize == 0))
	{
		return pOutputIdx;
	}
	TraversalBT(nums, numsSize, 0, target, &pOutputIdx[0], &pOutputIdx[1]);

	return pOutputIdx;
}