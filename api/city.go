package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nizaliyev7/employess/db/sqlc"
)

type createCityRequest struct {
	Name     string `json:"name" binding:"required,alphanum"`
	CityCode string `json:"city_code" binding:"required,alphanum"`
	IsActive bool   `json:"is_active"`
}

type CitiesResponse struct {
	Name      string    `json:"name"`
	CityCode  string    `json:"city_code"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func newCityResponse(city db.Cities) CitiesResponse {
	return CitiesResponse{
		Name:      city.Name,
		CityCode:  city.CityCode,
		IsActive:  city.IsActive,
		CreatedAt: city.CreatedAt,
	}
}
func (server *Server) createCity(ctx *gin.Context) {
	var req createCityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCityParams{
		Name:     req.Name,
		CityCode: req.CityCode,
		IsActive: req.IsActive,
	}

	city, err := server.store.CreateCity(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newCityResponse(city)
	ctx.JSON(http.StatusOK, rsp)
}

type updateCitiesRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateCitiesBodyRequest struct {
	Name     string `json:"name"`
	CityCode string `json:"city_code"`
	IsActive bool   `json:"is_active"`
}

func (server *Server) updateCities(ctx *gin.Context) {
	var reqID updateCitiesRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqBody updateCitiesBodyRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCityParams{
		ID:       reqID.ID,
		Name:     reqBody.Name,
		CityCode: reqBody.CityCode,
		IsActive: reqBody.IsActive,
	}

	city, err := server.store.UpdateCity(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, city)
}

func (server *Server) getCities(ctx *gin.Context) {
	cities, err := server.store.GetEmployees(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cities)
}
