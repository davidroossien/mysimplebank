package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
// returns the Maker interface, which means our code must implement the token maker interface
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	// returns a function pointer
	return &JWTMaker{secretKey}, nil
}

// maker *JWTMaker is a "pointer receiver"
// https://go.dev/tour/methods/4
// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// here we access the JWTMaker struct variable secretKey, using the "pointer receiver"
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// maker *JWTMaker is a "pointer receiver"
// https://go.dev/tour/methods/4
// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// very important, not to let the caller specify their own signing method
	// so we explicitly check the type
	// see token.go type Keyfunc
	// Parse methods use this callback function to supply
	// the key for verification.  The function receives the parsed,
	// but unverified Token.  This allows you to use properties in the
	// Header of the token (such as `kid`) to identify which key to use.
	// type Keyfunc func(*Token) (interface{}, error)
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	// see parser.go, token.Claims.Valid() code
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		// convert the error
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	// make sure Claims matches our Payload struct type
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
