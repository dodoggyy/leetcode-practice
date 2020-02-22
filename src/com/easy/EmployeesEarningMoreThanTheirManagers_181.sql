# 2020年2月23日
# Use inner join:
# Runtime: 457 ms, faster than 78.98%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT E2.Name AS Employee
FROM Employee AS E1 INNER JOIN Employee AS E2
ON E1.Id = E2.ManagerId
WHERE E2.Salary > E1.Salary;