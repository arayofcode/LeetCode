SELECT contest_id, ROUND(COUNT(user_id) * 100.0 /(SELECT COUNT(user_id) FROM Users), 2) percentage
FROM Register INNER JOIN Users USING(user_id)
GROUP BY contest_id
ORDER BY percentage DESC, contest_id