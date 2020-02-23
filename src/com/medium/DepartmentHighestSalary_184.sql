# 2020年2月23日
# Use JOIN:
# Runtime: 845 ms, faster than 48.04%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT D1.Name AS 'Department', E1.Name AS 'Employee', Salary
FROM Employee E1 JOIN Department D1
ON E1.DepartmentId = D1.Id
WHERE (E1.DepartmentId, Salary) IN
    ( SELECT
        DepartmentId, MAX(Salary)
      FROM
        Employee
      GROUP BY DepartmentId
    )
;