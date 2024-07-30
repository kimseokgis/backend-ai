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

// getPasetoSecretKey generates a PASETO secret key from a given hex string.
// Returns the secret key or an error if key generation fails.
func getPasetoSecretKey(privatekey string) (*paseto.V4AsymmetricSecretKey, error) {
	key, err := paseto.NewV4AsymmetricSecretKeyFromHex(privatekey)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// signPasetoToken signs the token with the given secret key.
// Returns the signed token as a string.
func signPasetoToken(token *paseto.Token, key *paseto.V4AsymmetricSecretKey) string {
	return token.V4Sign(key, nil)
}