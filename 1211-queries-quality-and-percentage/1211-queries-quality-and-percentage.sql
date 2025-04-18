SELECT 
    query_name, 
    ROUND(AVG(rating * 1.0 / position), 2) quality, 
    ROUND(SUM((rating < 3)::int) * 100.0 / COUNT(*), 2) poor_query_percentage
FROM Queries
GROUP BY query_name