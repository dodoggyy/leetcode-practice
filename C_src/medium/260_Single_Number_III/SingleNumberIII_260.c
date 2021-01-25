#include <stdio.h>
#include <stdlib.h>

// Use bit manipulation
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 98.73%
// Memory Usage: 8.3 MB, less than 100.00%
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *singleNumber(int *nums, int numsSize, int *returnSize)
{
    int tmpXOR = 0;
    int i;
    // int *result = calloc(*returnSize = 2, sizeof(int));
    int *result = (int *)malloc(sizeof(int) * 2);
    memset(result, 0, sizeof(int) * 2);
    *returnSize = 2;
    for (i = 0; i < numsSize; i++)
    {
        tmpXOR ^= nums[i];
    }
    int tmpDiffBit = (tmpXOR & (-tmpXOR));
    for (i = 0; i < numsSize; i++)
    {
        if (tmpDiffBit & nums[i])
        {
            result[0] ^= nums[i];
        }
        else
        {
            result[1] ^= nums[i];
        }
    }
    return result;
}