#include <stdio.h>
#include <stdlib.h>

// Use bit manipulation
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 20 ms, faster than 19.21%
// Memory Usage: 8.1 MB, less than 63.64%
int singleNumber(int *nums, int numsSize)
{
    int i;
    int result = 0;
    for (i = 0; i < numsSize; i++)
    {
        result ^= nums[i];
    }
    return result;
}