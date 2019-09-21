# 2019年9月22日
# Use XOR operation
# Runtime: 140 ms, faster than 85.45%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
UPDATE salary
SET sex = CHAR(ASCII('f') ^ ASCII('m') ^ ASCII(sex));


# 2019年9月22日
# Use CASE WHEN THEN condition
# Runtime: 142 ms, faster than 81.27%
# Memory Usage: 0B, less than 100.00%
# Write your MySQL query statement below
UPDATE salary
SET sex = CASE sex
        WHEN 'm' THEN 'f'
        ELSE 'm'
    END;