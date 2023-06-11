package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// reduces verbosity of console output for test results
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
