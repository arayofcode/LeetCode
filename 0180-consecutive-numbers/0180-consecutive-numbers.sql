SELECT DISTINCT l1.num ConsecutiveNums
FROM Logs l1 LEFT JOIN Logs l2 ON l2.id = l1.id + 1 LEFT JOIN Logs l3 ON l3.id = l2.id + 1
WHERE l1.num = l2.num AND l2.num = l3.num;