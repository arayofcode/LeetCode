SELECT DISTINCT Department.name "Department", Employee.name "Employee", Employee.salary "Salary"
FROM Employee LEFT JOIN Department ON Employee.departmentId = Department.id
GROUP BY Department.name, Employee.departmentId, Employee.name, Employee.salary
HAVING Employee.salary IN (
    SELECT DISTINCT salary
    FROM Employee e1
    WHERE e1.departmentId = Employee.departmentId
    GROUP BY salary
    ORDER BY salary DESC
    LIMIT 3
);