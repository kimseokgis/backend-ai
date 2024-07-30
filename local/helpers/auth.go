package helpers

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

// createPasetoToken initializes a new PASETO token with the given username.
// Sets the issued at, not before, and expiration times.
func createPasetoToken(username string) *paseto.Token {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.SetString("user", username)
	return token
}
