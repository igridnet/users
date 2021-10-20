package users

import (
	"fmt"
	"time"
)

//Randomizer Rand Generator generates a random string of specified length n
//It is essential in creating reference ids and tokens
type Randomizer interface {
	Get(length int) (value string)
}

// Tokenizer specifies API for encoding and decoding between string and Key.
type Tokenizer interface {
	// Issue converts API Key to its string representation.
	Issue(Key) (string, error)

	// Parse extracts API Key data from string token.
	Parse(string) (Key, error)
}


// Hasher specifies an API for generating hashes of an arbitrary textual
// content.
type Hasher interface {
	// Hash generates the hashed string from plain-text.
	Hash(plainText string) (hashedPassword string, err error)

	// Compare compares plain-text version to the hashed one. An error should
	// indicate failed comparison.
	Compare(hashedPassword string, plainText string) error
}

//Key ...
//iss (issuer): Issuer of the JWT
//sub (subject): Subject of the JWT (the user)
//aud (audience): Recipient for which the JWT is intended
//exp (expiration time): Time after which the JWT expires
//nbf (not before time): Time before which the JWT must
//not be accepted for processing
//iat (issued at time): Time at which the JWT was issued;
//can be used to determine age of the JWT
type Key struct {
	Issuer    string    `json:"iss"`
	Purpose   string    `json:"purpose"` //api for things, access and refresh for users,
	Subject   string    `json:"sub"`     //userid
	Audience  string    `json:"aud"`     //igrid-message-bus, igrid-user-services
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

func NewKey(id, purpose string) Key {
	return Key{
		Purpose:   purpose,
		Subject:   id,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

func (key Key) String() string {
	return fmt.Sprintf("iss: %v, purpose: %v, sub: %v aud: %v, iat: %v, exp: %v\n",
		key.Issuer, key.Purpose, key.Subject, key.Audience, key.IssuedAt, key.ExpiresAt)
}

// IDProvider specifies an API for generating unique identifiers.
type IDProvider interface {
	// ID generates the unique identifier.
	ID() (string, error)
}
