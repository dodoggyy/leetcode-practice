# 2020年2月24日
# Using sub-query and LIMIT:
# Runtime: 156 ms, faster than 89.14%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT (
SELECT DISTINCT Salary
FROM Employee
ORDER BY Salary DESC LIMIT 1 OFFSET 1
) AS SecondHighestSalary;