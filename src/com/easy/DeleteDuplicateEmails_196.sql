# 2020年2月22日
# Use simple query:
# Runtime: 659 ms, faster than 97.07%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
DELETE p1 
FROM Person p1, Person p2
WHERE p1.Email = p2.Email AND p1.Id > p2.Id;


# 2020年2月22日
# Use sub query:
# Runtime: 1412 ms, faster than 36.24%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
DELETE
FROM Person
WHERE id NOT IN (
   SELECT id
   FROM (
        SELECT MIN(id) AS id
        FROM Person
        GROUP BY Email
        ) t
   );