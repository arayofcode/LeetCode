SELECT w1.id
FROM Weather w1 JOIN Weather w2
    ON w2.recordDate = w1.recordDate - interval '1 day'
WHERE w1.temperature > w2.temperature;