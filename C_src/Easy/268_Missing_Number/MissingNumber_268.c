#include <stdio.h>
#include <stdlib.h>

// Use bit manipulation
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 95.72%
// Memory Usage: 7.9 MB, less than 71.43%
int missingNumber(int *nums, int numsSize)
{
    int result = numsSize;
    int i;
    for (i = 0; i < numsSize; i++)
    {
        result ^= i;
        result ^= nums[i];
    }
    return result;
}