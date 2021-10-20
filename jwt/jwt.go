package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/igridnet/users"
)

var _ users.Tokenizer = (*tokenizer)(nil)

type (
	tokenizer struct {
		signingKey []byte
	}

	claims struct {
		Purpose string `json:"purpose"`
		jwt.StandardClaims
	}
)

func keyToClaims(key users.Key) claims {
	return claims{
		Purpose: key.Purpose,
		StandardClaims: jwt.StandardClaims{
			Audience:  key.Audience,
			ExpiresAt: key.ExpiresAt,
			IssuedAt:  key.IssuedAt,
			Issuer:    key.Issuer,
			NotBefore: key.NotBefore,
			Subject:   key.Subject,
		},
	}
}

func claimsToKey(c claims) users.Key {

	return users.Key{
		Purpose: c.Purpose,
		Audience:  c.Audience,
		Issuer:    c.Issuer,
		Subject:   c.Subject,
		IssuedAt:  c.IssuedAt,
		NotBefore: c.NotBefore,
		ExpiresAt: c.ExpiresAt,
	}
}

func NewTokenizer(secret string) users.Tokenizer {
	return &tokenizer{signingKey: []byte(secret)}
}

func (t *tokenizer) Issue(key users.Key) (string, error) {
	c := keyToClaims(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString(t.signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (t *tokenizer) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return t.signingKey, nil
}

func (t *tokenizer) Parse(tokenString string) (users.Key, error) {
	cl := new(claims)
	token, err := jwt.ParseWithClaims(tokenString, cl, t.keyFunc)

	if err != nil {
		return users.Key{}, fmt.Errorf("could not parse token %w", err)
	}
	tkn, ok := token.Claims.(*claims)
	if !ok {
		return users.Key{}, fmt.Errorf("could not parse token to claims")
	}
	if err := tkn.Valid(); err != nil {
		return users.Key{}, fmt.Errorf("could not validate parsed token %w", err)
	}
	fmt.Printf("%+v", tkn)
	return claimsToKey(*tkn), nil
}
