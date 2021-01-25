#include <stdio.h>
#include <stdlib.h>

// Use bit manipulation
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 24 ms, faster than 95.35%
// Memory Usage: 9.2 MB, less than 100.00%
#define MAX(a, b) ((a) > (b) ? (a) : (b))

int maxProduct(char **words, int wordsSize)
{
    int *u32Value = (int *)calloc(wordsSize, sizeof(int));
    int u32Result = 0;
    for (int i = 0; i < wordsSize; i++)
    {
        for (char *word = words[i]; *word; word++)
        {
            u32Value[i] |= 1 << (*word - 'a');
        }
    }
    for (int i = 0; i < wordsSize; i++)
    {
        for (int j = i + 1; j < wordsSize; j++)
        {
            if ((u32Value[i] & u32Value[j]) == 0)
            {
                u32Result = MAX(u32Result, strlen(words[i]) * strlen(words[j]));
            }
        }
    }
    return u32Result;
}
