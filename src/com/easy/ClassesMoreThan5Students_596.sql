# 2019年9月25日
# Use simple query solution
# Runtime: 465 ms, faster than 5.15%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT class
FROM 
    (   SELECT class, COUNT(DISTINCT student) AS num
        FROM courses
        GROUP BY class) AS tmp_table
WHERE num >=5;


# 2019年9月25日
# Use simple query solution 2
# Runtime: 642 ms, faster than 5.15%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT class
FROM courses
GROUP BY class
HAVING COUNT(DISTINCT student) >=5;