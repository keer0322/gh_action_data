SELECT 
    age,
    SUM(CASE WHEN diagnosis_code = 'diabetes' THEN 1 ELSE 0 END) AS diabetes_count,
    COUNT(*) AS total_patients,
    SUM(CASE WHEN diagnosis_code = 'diabetes' THEN 1 ELSE 0 END) / COUNT(*) AS diabetes_prevalence,
    (SUM(CASE WHEN diagnosis_code = 'diabetes' THEN 1 ELSE 0 END) / COUNT(*)) / (1 - (SUM(CASE WHEN diagnosis_code = 'diabetes' THEN 1 ELSE 0 END) / COUNT(*))) AS likelihood_ratio
FROM 
    patients
GROUP BY 
    age
ORDER BY 
    age;
