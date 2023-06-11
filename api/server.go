package api

import (
	db "github.com/davidroossien/mysimplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// (route, handler) could also place middleware here
	// handler must accept gin context
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// returns a gin.H map of strings
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
