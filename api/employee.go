package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/nizaliyev7/employess/db/sqlc"
)

type createEmployeeRequest struct {
	FirstName  string `json:"first_name" binding:"required,alphanum"`
	LastName   string `json:"last_name" binding:"required,alphanum"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone" binding:"required"`
	CityID     int64  `json:"city_id" binding:"required"`
}

type EmployeeResponse struct {
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
	Phone      string    `json:"phone"`
	CityID     int64     `json:"city_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func newEmployeeResponse(employee db.Employees) EmployeeResponse {
	return EmployeeResponse{
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		MiddleName: employee.MiddleName,
		Phone:      employee.Phone,
		CityID:     employee.CityID,
		CreatedAt:  employee.CreatedAt,
	}
}

func (server *Server) createEmployee(ctx *gin.Context) {
	var req createEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEmployeeParams{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
		Phone:      req.Phone,
		CityID:     req.CityID,
	}

	employee, err := server.store.CreateEmployee(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newEmployeeResponse(employee)
	ctx.JSON(http.StatusOK, rsp)
}

type getEmployeeRequest struct {
	ID int64 `form:"id" binding:"required,min=1"`
}

func (server *Server) getEmployee(ctx *gin.Context) {
	var req getEmployeeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	employee, err := server.store.GetEmployee(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, employee)
}

func (server *Server) getEmployees(ctx *gin.Context) {
	employees, err := server.store.GetEmployees(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, employees)
}

type updateEmployeesRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateEmployeeBodyRequest struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone" binding:"required"`
	CityID     int64  `json:"city_id" binding:"required"`
}

func (server *Server) updateEmployees(ctx *gin.Context) {
	var reqID updateEmployeesRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqBody updateEmployeeBodyRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateEmployeeParams{
		ID:         reqID.ID,
		FirstName:  reqBody.FirstName,
		LastName:   reqBody.LastName,
		MiddleName: reqBody.MiddleName,
		Phone:      reqBody.Phone,
		CityID:     reqBody.CityID,
	}

	employee, err := server.store.UpdateEmployee(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}
