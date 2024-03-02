package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/saurabhsingh121/simplebank/db/sqlc"
)

// Server serves HTTP request to our banking service
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing
func NewServer(db db.Store) *Server{
  server := &Server{store: db}
  router := gin.Default()

  // add routes to the server
  router.POST("/accounts", server.createAccount)
  router.GET("/accounts/:id", server.getAccount)
  router.GET("/accounts", server.listAccount)
  server.router = router
  return server
}

// Start run a HTTP server on a specific address
func(server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error)gin.H{
	return gin.H{"error": err.Error()}
}