// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateCity(ctx context.Context, arg CreateCityParams) (Cities, error)
	CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employees, error)
	GetCities(ctx context.Context) ([]Cities, error)
	GetEmployees(ctx context.Context) ([]GetEmployeesRow, error)
	UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (Employees, error)
}

var _ Querier = (*Queries)(nil)
