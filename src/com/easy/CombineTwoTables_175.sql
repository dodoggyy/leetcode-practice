# 2020年2月22日
# Use left join:
# Runtime: 661 ms, faster than 22.30%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT FirstName, LastName, City, State
FROM Person AS P LEFT JOIN Address AS A
ON P.PersonId = A.PersonId;