SELECT Employee.name "Employee"
FROM Employee INNER JOIN Employee Manager
    ON Employee.managerId = Manager.id AND Employee.salary > Manager.salary;