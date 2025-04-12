SELECT
    student_id,
    student_name,
    subject_name,
    COUNT(Examinations.subject_name) "attended_exams"
FROM (SELECT * FROM Students, Subjects)
    LEFT JOIN Examinations USING(student_id, subject_name)
GROUP BY subject_name, student_id, student_name
ORDER BY student_id, subject_name;