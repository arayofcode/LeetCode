SELECT machine_id, ROUND(AVG(a2.timestamp - a1.timestamp)::numeric, 3) "processing_time" 
FROM Activity a1 INNER JOIN Activity a2 USING(machine_id, process_id)
WHERE a1.activity_type = 'start' AND a2.activity_type = 'end'
GROUP BY machine_id