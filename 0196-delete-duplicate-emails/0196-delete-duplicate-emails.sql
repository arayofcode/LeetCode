WITH cte AS (
  SELECT id, ROW_NUMBER() OVER (PARTITION BY email ORDER BY id) AS row_num
  FROM Person
)
DELETE FROM Person
WHERE id IN (SELECT id FROM cte WHERE row_num > 1);
-- DELETE FROM Person p1 USING Person p2
-- WHERE p1.email = p2.email AND p1.id > p2.id;

-- DELETE FROM Person
-- WHERE id NOT IN (
--     SELECT MIN(id) FROM Person GROUP BY email
-- );