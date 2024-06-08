package randomindex

import (
	"math/rand"
)

type Rand struct {
	Random *rand.Rand
}

func New(seed int64) *Rand {
	src := rand.NewSource(seed)
	random := rand.New(src)
	return &Rand{Random: random}
}
