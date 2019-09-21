# 2019年9月22日
# Use Simple solution
# Runtime: 237 ms, faster than 48.72%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
SELECT name, population, area
FROM World
WHERE area > 3000000 OR population > 25000000;