#include <stdio.h>
#include <stdlib.h>

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 43.48%
// Memory Usage: 6.8 MB, less than 100.00%
char *toHex(int num)
{
    char hex[] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
    unsigned int n = num;
    char *str = calloc(9, sizeof(char));
    char *result = str + 8;

    while (str != result)
    {
        *(--result) = hex[(n & 0xf)];
        n >>= 4;
    }

    // move to first digit
    while (*result == '0' && *(result + 1) != '\0')
    {
        result++;
    }

    return result;
}