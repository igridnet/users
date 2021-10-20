package hasher

import (
	"github.com/igridnet/users"
	"golang.org/x/crypto/bcrypt"
)

var (
	_ users.Hasher = (*hasher)(nil)
)

type (
	hasher struct {}
)

func New()users.Hasher{
	return &hasher{}
}

func (h *hasher) Hash(plainText string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	return string(bytes), err
}

func (h *hasher) Compare(hashedPassword string, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
}


