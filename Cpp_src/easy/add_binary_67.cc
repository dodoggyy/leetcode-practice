#include <algorithm>
#include <iostream>
using namespace std;

// Use sequential parsing
// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.6 MB, less than 87.27%
class Solution
{
public:
    string addBinary(string a, string b)
    {
        string result = "";
        int length_a = a.size(), length_b = b.size();
        int length = max(length_a, length_b);

        reverse(a.begin(), a.end());
        reverse(b.begin(), b.end());

        int carry = 0;

        for (int i = 0; i < length; i++)
        {
            int tmp_a = (i < length_a) ? a[i] - '0' : 0;
            int tmp_b = (i < length_b) ? b[i] - '0' : 0;
            int val = tmp_a + tmp_b + carry;
            carry = val / 2;

            result.insert(result.begin(), (val % 2) + '0');
        }
        if (carry == 1)
        {
            result.insert(result.begin(), '1');
        }

        return result;
    }
};

// Use sequential parsing without reverse
// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 4 ms, faster than 80.35%
// Memory Usage: 8.7 MB, less than 76.36%
class Solution2
{
public:
    string addBinary(string a, string b)
    {
        string result = "";
        int length_a = a.size(), length_b = b.size();
        int length = max(length_a, length_b);

        int carry = 0;

        for (int i = 0; i < length; i++)
        {
            int tmp_a = (i < length_a) ? a[length_a - i - 1] - '0' : 0;
            int tmp_b = (i < length_b) ? b[length_b - i - 1] - '0' : 0;
            int val = tmp_a + tmp_b + carry;
            carry = val / 2;

            result.insert(result.begin(), (val % 2) + '0');
        }
        if (carry == 1)
        {
            result.insert(result.begin(), '1');
        }

        return result;
    }
};