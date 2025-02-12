// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: employees.sql

package db

import (
	"context"
	"time"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employees (first_name, last_name, middle_name, phone, city_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, first_name, last_name, middle_name, phone, created_at, city_id
`

type CreateEmployeeParams struct {
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	MiddleName string         `json:"middle_name"`
	Phone      string         `json:"phone"`
	CityID     int64          `json:"city_id"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employees, error) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.Phone,
		arg.CityID,
	)
	var i Employees
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.Phone,
		&i.CreatedAt,
		&i.CityID,
	)
	return i, err
}

const getEmployee = `-- name: GetEmployee :one
SELECT id, first_name, last_name, middle_name, phone, created_at, city_id FROM employees
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEmployee(ctx context.Context, id int64) (Employees, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, id)
	var i Employees
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.Phone,
		&i.CreatedAt,
		&i.CityID,
	)
	return i, err
}

const getEmployees = `-- name: GetEmployees :many
SELECT employees.id, employees.first_name, employees.last_name, employees.middle_name, 
       employees.phone, employees.created_at, cities.name AS city
FROM employees
JOIN cities ON employees.city_id = cities.id
ORDER BY employees.created_at DESC
`

type GetEmployeesRow struct {
	ID         int64          `json:"id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	MiddleName string         `json:"middle_name"`
	Phone      string         `json:"phone"`
	CreatedAt  time.Time      `json:"created_at"`
	City       string         `json:"city"`
}

func (q *Queries) GetEmployees(ctx context.Context) ([]GetEmployeesRow, error) {
	rows, err := q.db.QueryContext(ctx, getEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEmployeesRow{}
	for rows.Next() {
		var i GetEmployeesRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.MiddleName,
			&i.Phone,
			&i.CreatedAt,
			&i.City,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmployee = `-- name: UpdateEmployee :one
UPDATE employees 
SET first_name = $2, last_name = $3, middle_name = $4, phone = $5, city_id = $6 
WHERE id = $1 
RETURNING id, first_name, last_name, middle_name, phone, created_at, city_id
`

type UpdateEmployeeParams struct {
	ID         int64          `json:"id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	MiddleName string         `json:"middle_name"`
	Phone      string         `json:"phone"`
	CityID     int64          `json:"city_id"`
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (Employees, error) {
	row := q.db.QueryRowContext(ctx, updateEmployee,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.Phone,
		arg.CityID,
	)
	var i Employees
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.Phone,
		&i.CreatedAt,
		&i.CityID,
	)
	return i, err
}
