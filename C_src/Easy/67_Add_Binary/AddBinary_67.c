#include <stdio.h>
#include <stdlib.h>

// Use sequential parsing:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 62.18%
// Memory Usage: 6.9 MB, less than 100.00%
char *addBinary(char *a, char *b)
{
    int lengthA = strlen(a);
    int lengthB = strlen(b);
    int lengthResult = lengthA > lengthB ? lengthA : lengthB;
    char *result = (char *)malloc(lengthResult + 2);

    int carry = 0, lengthTmp = lengthResult;

    result[lengthResult + 1] = '\0';
    for (int i = 0; i < lengthResult; i++)
    {
        int tmpA = (i < lengthA) ? a[lengthA - i - 1] - '0' : 0;
        int tmpB = (i < lengthB) ? b[lengthB - i - 1] - '0' : 0;
        int val = tmpA + tmpB + carry;
        result[lengthTmp - i] = (val & 1) + '0';
        carry = (val >> 1);
    }
    result[0] = carry + '0';

    // The (carry^1) will determine if we need to skip the first character (no overflow, i.e. last carry = 0), or not overflow, last carry = 1
    return result + (carry ^ 1);
}