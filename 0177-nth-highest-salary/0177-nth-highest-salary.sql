CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
    IF N <= 0 THEN
        RETURN QUERY SELECT NULL::int;
        RETURN;
    END IF;
    
    RETURN QUERY (
        WITH ranked_salaries AS (
            SELECT DISTINCT Employee.salary
            FROM Employee
            ORDER BY Employee.salary DESC
            LIMIT N
        )
        SELECT ranked_salaries.salary
        FROM ranked_salaries
        ORDER BY ranked_salaries.salary DESC
        LIMIT 1
        OFFSET N-1
    );
    
    IF NOT FOUND THEN
        RETURN QUERY SELECT NULL::int;
    END IF;
END;
$$ LANGUAGE plpgsql;