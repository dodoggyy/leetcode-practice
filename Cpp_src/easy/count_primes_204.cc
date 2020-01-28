#include <algorithm>
#include <iostream>
using namespace std;

// Use simple judgment under sqrt(n)
// Time Complexity: O(n^1.5)
// Space Complexity: O(1)
// Runtime: 844 ms, faster than 5.37%
// Memory Usage: 8.2 MB, less than 91.67%
class Solution
{
public:
    int countPrimes(int n)
    {
        if (n <= 2)
        {
            return 0;
        }
        int result = 1; // start from 3, and 2 is prime

        for (int i = 3; i < n; i++)
        {
            if (bIsPrime(i))
            {
                result++;
            }
        }

        return result;
    }

    bool bIsPrime(int n)
    {
        if (n % 2 == 0)
        {
            return false;
        }
        for (long i = 3; i * i <= n; i += 2)
        {
            if (n % i == 0)
            {
                return false;
            }
        }
        return true;
    }
};

// Use sieve of Eratosthenes
// Time Complexity: O(nloglogn)
// Space Complexity: O(n)
// Runtime: 28 ms, faster than 85.17%
// Memory Usage: 9.8 MB, less than 16.67%
class Solution2
{
public:
    int countPrimes(int n)
    {
        if (n <= 2)
        {
            return 0;
        }
        int result = 0;
        bool b_prime[n];
        for (int i = 0; i < n; i++)
        {
            b_prime[i] = true;
        }
        b_prime[0] = false;
        b_prime[1] = false;
        for (int i = 2; i < n; i++)
        {
            if (!b_prime[i])
            {
                continue;
            }
            result++;
            for (int j = 2 * i; j < n; j += i)
            {
                b_prime[j] = false;
            }
        }

        return result;
    }
};