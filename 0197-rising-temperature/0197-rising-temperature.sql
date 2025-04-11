SELECT w2.id "Id"
FROM Weather w2 LEFT JOIN Weather w1
    ON w2.recordDate = w1.recordDate + INTERVAL '1 day'
WHERE w2.temperature > w1.temperature