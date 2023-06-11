package api

import (
	"github.com/gin-gonic/gin"
	// db "github.com/techschool/simplebank/db/sqlc"
	db "github.com/davidroossien/mysimplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
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
