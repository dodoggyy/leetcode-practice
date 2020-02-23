# 2020年2月23日
# Use left join:
# Runtime: 283 ms, faster than 84.22% 
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT Name AS Customers
FROM Customers C1 LEFT JOIN Orders O1
ON C1.Id = O1.CustomerId 
WHERE CustomerId IS NULL;


# 2020年2月23日
# Use sub-query:
# Runtime: 213 ms, faster than 99.41%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT Name AS 'Customers'
FROM Customers
WHERE 
    Id NOT IN (
    SELECT CustomerId
    FROM Orders 
    );