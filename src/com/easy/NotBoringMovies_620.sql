# 2019年9月25日
# Use simple query solution
# Runtime: 454 ms, faster than 5.15%
# Memory Usage: 0B, less than 100.00%
# ` ` can be ignore
# Write your MySQL query statement below
SELECT *
FROM cinema
WHERE `id`%2 = 1 AND `description` != "boring"
ORDER BY `rating` DESC;