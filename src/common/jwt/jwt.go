package jwt

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"github.com/distanceNing/testapp/src/common/crypto"
	"github.com/google/uuid"
	"io"
	"strings"
	"time"
)

type Token struct {
	Exp int64
	T   string
}

// AppleJWTGenerator
type AppleJWTGenerator struct {
	tokens  map[string]*Token
	g       *JWTGenerator
	timeout int64
}

func NewAppleJWTGenerator() *AppleJWTGenerator {
	return &AppleJWTGenerator{g: NewJWTGenerator(), timeout: 60 * 60, tokens: make(map[string]*Token)}
}

type JWTParams struct {
	AppleId    string
	Kid        string
	PrivateKey string
	Bid        string
	IssuerId   string
}

func (ag *AppleJWTGenerator) Get(params *JWTParams) (string, error) {
	if len(params.Kid) == 0 || len(params.Bid) == 0 || len(params.AppleId) == 0 || len(params.IssuerId) == 0 {
		return "", errors.New("params is empty")
	}
	type AppleJwtHeader struct {
		Alg string `json:"alg,omitempty"`
		Kid string `json:"kid,omitempty"`
		Typ string `json:"typ,omitempty"`
	}

	type AppleJwtPayload struct {
		Iss   string `json:"iss,omitempty"`
		Aud   string `json:"aud,omitempty"`
		Nonce string `json:"nonce,omitempty"`
		Bid   string `json:"bid,omitempty"`
		Iat   int64  `json:"iat,omitempty"`
		Exp   int64  `json:"exp,omitempty"`
	}
	key := params.AppleId + params.Bid
	val, ok := ag.tokens[key]
	if ok {
		return val.T, nil
	}
	h := &AppleJwtHeader{"ES256", params.Kid, "JWT"}
	now := time.Now().Unix()
	exp := now + ag.timeout
	u := uuid.New()
	p := &AppleJwtPayload{params.IssuerId, "appstoreconnect-v1", u.String(), params.Bid, now, exp}
	token, err := ag.g.Get(h, p, strings.NewReader(params.PrivateKey))
	if err != nil {
		return "", err
	}
	ag.tokens[key] = &Token{p.Exp, token}
	return token, nil
}

type JWTGenerator struct {
}

func NewJWTGenerator() *JWTGenerator {
	return &JWTGenerator{}
}

func (g *JWTGenerator) Get(header interface{}, payload interface{}, priKey io.Reader) (string, error) {
	headerJson, err := json.Marshal(header) //转换成JSON返回的是byte[]
	if err != nil {
		return "", err
	}
	payloadJson, err1 := json.Marshal(payload) //转换成JSON返回的是byte[]
	if err1 != nil {
		return "", err1
	}

	encodeHeader := crypto.EncodeSegment(headerJson)
	encodePayload := crypto.EncodeSegment(payloadJson)
	signString := encodeHeader + "." + encodePayload
	var prk *ecdsa.PrivateKey
	prk, err = crypto.ParseECPrivateKeyFromPEM(priKey)
	if err != nil {
		return "", err
	}
	sign, err := crypto.SigningMethodES256.Sign(signString, prk)
	if err != nil {
		return "", err
	}
	return signString + "." + sign, nil
}
