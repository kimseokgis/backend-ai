package helpers

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

func GenerateToken(username, privatekey string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))