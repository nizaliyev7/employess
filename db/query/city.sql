-- name: CreateCity :one
INSERT INTO cities (name,city_code,is_active) 
VALUES ($1,$2,$3) 
RETURNING *;

-- name: GetCities :many
SELECT * FROM cities;

-- name: UpdateCity :one
UPDATE cities
SET name = $2,
    city_code = $3,
    is_active = $4
WHERE id = $1
RETURNING id, name, city_code, is_active, created_at;