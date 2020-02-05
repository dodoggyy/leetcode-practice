#include <stdio.h>
#include <stdlib.h>

// Use 1 parameter and result array record:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 56 ms, faster than 28.28%
// Memory Usage: 15.3 MB, less than 100.00%
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *productExceptSelf(int *nums, int numsSize, int *returnSize)
{
    int *result = calloc(sizeof(int), *returnSize = numsSize);

    *result = 1;
    for (int i = 1; i < numsSize; i++)
    {
        *(result + i) = *(result + i - 1) * *(nums + i - 1);
    }

    int R = 1;
    for (int i = numsSize - 1; i >= 0; i--)
    {
        *(result + i) *= R;
        R *= *(nums + i);
    }

    return result;
}

int *productExceptSelf2(int *nums, int numsSize, int *returnSize)
{
    int *result = calloc(sizeof(int), *returnSize = numsSize);

    result[0] = 1;
    for (int i = 1; i < numsSize; i++)
    {
        result[i] = result[i - 1] * nums[i - 1];
    }

    int R = 1;
    for (int i = numsSize - 1; i >= 0; i--)
    {
        result[i] *= R;
        R *= nums[i];
    }

    return result;
}