package helpers

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

// createPasetoToken initializes a new PASETO token with the given username.


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