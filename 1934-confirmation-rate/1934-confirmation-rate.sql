SELECT user_id, COALESCE(ROUND(AVG(CASE WHEN action = 'confirmed' THEN 1 ELSE 0 END), 2), 0.0) "confirmation_rate"
FROM Signups LEFT JOIN Confirmations USING(user_id)
GROUP BY user_id