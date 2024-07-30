package helpers

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

// createPasetoToken initializes a new PASETO token with the given username.


// getPasetoSecretKey generates a PASETO secret key from a given hex string.


// signPasetoToken signs the token with the given secret key.


// GenerateToken generates a PASETO token for the given username and private key.
// Returns the signed token or an error if any step fails.
func GenerateToken(username, privatekey string) (string, error) {
	token := createPasetoToken(username)
	key, err := getPasetoSecretKey(privatekey)
	if err != nil {
		return "", err
	}
	return signPasetoToken(token, key), nil
}