#include <stdio.h>
#include <stdlib.h>

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 6.7 MB, less than 100.00%
char *convertToBase7(int num)
{
    char *result = calloc(11, sizeof(char));
    if (num == 0)
    {
        return "0";
    }

    if (num < 0)
    {
        sprintf(result, "-%s", convertToBase7(abs(num)));
    }
    else
    {
        char *tmp = convertToBase7(num / 7);
        if (tmp == "0")
        {
            tmp = "";
        }
        sprintf(result, "%s%d", tmp, num % 7);
    }
    return result;
}
