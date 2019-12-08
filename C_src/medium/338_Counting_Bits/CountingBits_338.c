#include <stdio.h>
#include <stdlib.h>

// Use bit manipulation:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 40 ms, faster than 60.94%
// Memory Usage: 11.6 MB, less than 100.00%
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *countBits(int num, int *returnSize)
{
    int *pu32Result = calloc(*returnSize = num + 1, sizeof(int));
    for (int i = 1; i <= num; i++)
    {
        pu32Result[i] = pu32Result[i & (i - 1)] + 1;
    }

    return pu32Result;
}
