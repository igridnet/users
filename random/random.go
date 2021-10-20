package random

import (
	"github.com/igridnet/users"
	"math/rand"
	"time"
)

var (
	_ users.Randomizer = (*randomizer)(nil)
)

type randomizer struct {
	pool []rune
}

//Get returns a random string of a specified length
//credit to https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986
func (r *randomizer) Get(length int) (val string) {
	b := make([]rune, length)
	for i := range b {
		b[i] = r.pool[rand.Intn(len(r.pool))]
	}
	return string(b)
}

func New(pool []rune) users.Randomizer {
	rand.Seed(time.Now().UnixNano())
	return &randomizer{pool: pool}
}
