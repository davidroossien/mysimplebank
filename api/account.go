package api

import (
	"database/sql"
	"net/http"

	db "github.com/davidroossien/mysimplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		// sends a response
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	Owner    string `form:"owner" binding:"required"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		// sends a response
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Owner:  req.Owner,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

// gin uses go-playground/validator/v10
// https://pkg.go.dev/github.com/go-playground/validator
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// sends a response
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	// https://go.dev/tour/methods/15
	// Type assertions
	// A type assertion provides access to an interface value's underlying concrete value.
	// t := i.(T)
	// This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
	// If i does not hold a T, the statement will trigger a panic.
	// To test whether an interface value holds a specific type, a type assertion can return two values:
	// the underlying value and a boolean value that reports whether the assertion succeeded.
	// t, ok := i.(T)
	// If i holds a T, then t will be the underlying value and ok will be true.
	// If not, ok will be false and t will be the zero value of type T, and no panic occurs.
	// Note the similarity between this syntax and that of reading from a map.
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		// check if err is of type *pq.Error or not
		// if err is of type *pq.Error, then pqErr will be assigned to err and ok set true
		// then the ok section will be executed
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
