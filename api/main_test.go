package api

import (
	"os"
	"testing"
	"time"

	db "github.com/davidroossien/mysimplebank/db/sqlc"
	"github.com/davidroossien/mysimplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	// reduces verbosity of console output for test results
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
