package token

import "time"

// so we can switch between different token makers when we want to
// maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

// type Maker interface {
// 	// CreateToken creates a new token for a specific username and duration
// 	CreateToken(username string, duration time.Duration) (string, *Payload, error)

// 	// VerifyToken checks if the token is valid or not
// 	VerifyToken(token string) (*Payload, error)
// }
