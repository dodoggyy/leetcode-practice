# 2019年9月26日
# Use simple query solution
# Runtime: 510 ms, faster than 5.03%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT Email
FROM Person
GROUP BY Email
HAVING COUNT(*) >= 2;