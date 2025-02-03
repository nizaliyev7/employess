-- name: CreateCity :one
INSERT INTO cities (name,city_code,is_active) 
VALUES ($1,$2,$3) 
RETURNING *;

-- name: GetCities :many
SELECT * FROM cities;
