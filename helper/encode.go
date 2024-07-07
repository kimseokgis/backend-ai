package helper

import (
	"encoding/json"
	"fmt"

	"aidanwoods.dev/go-paseto"
	"github.com/kimseokgis/backend-ai/model"
)

func Decoder(publickey, tokenstr string) (payload model.Payload, err error) {
	var token *paseto.Token
	var pubKey paseto.V4AsymmetricPublicKey
	pubKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(publickey)
	if err != nil {
		fmt.Println("Decode NewV4AsymmetricPublicKeyFromHex : ", err)
	}
	parser := paseto.NewParser()
	token, err = parser.ParseV4Public(pubKey, tokenstr, nil)
	if err != nil {
		fmt.Println("Decode ParseV4Public : ", err)
	} else {
		json.Unmarshal(token.ClaimsJSON(), &payload)
	}
	return payload, err
}
