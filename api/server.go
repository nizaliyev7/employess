package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nizaliyev7/employess/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/employees", server.createEmployee)
	router.GET("/employee", server.getEmployee)
	router.GET("/employees", server.getEmployees)
	router.PUT("/employees/:id", server.updateEmployees)

	router.POST("/cities", server.createCity)
	router.PUT("/cities/:id", server.updateCities)
	router.GET("/cities", server.getCities)

	server.router = router
	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
