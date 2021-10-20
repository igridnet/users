package api

import (
	"github.com/go-pg/pg/v10"
	"github.com/igridnet/users"
	"github.com/igridnet/users/factory"
)

func NewClient(db *pg.DB, f *factory.Factory, tokenizer users.Tokenizer, hasher users.Hasher) *users.Client {
	return &users.Client{
		Tokenizer: tokenizer,
		Hasher:    hasher,
		Factory:   f,
		Db:        db,
	}
}
