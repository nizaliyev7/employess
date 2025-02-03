-- name: CreateEmployee :one
INSERT INTO employees (first_name, last_name, middle_name, phone, city_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetEmployees :many
SELECT employees.id, employees.first_name, employees.last_name, employees.middle_name, 
       employees.phone, employees.created_at, cities.name AS city
FROM employees
JOIN cities ON employees.city_id = cities.id
ORDER BY employees.created_at DESC;

-- name: UpdateEmployee :one
UPDATE employees 
SET first_name = $2, last_name = $3, middle_name = $4, phone = $5, city_id = $6 
WHERE id = $1 
RETURNING *;