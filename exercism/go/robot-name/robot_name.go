package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.

var names = map[string]bool{}

type Robot struct {
	name string
}

func genName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%s%d%d%d", string('A'+rune(rand.Intn(26))), string('A'+rune(rand.Intn(26))), rand.Intn(10), rand.Intn(10), rand.Intn(10))
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	r.name = genName()
	for names[r.name] {
		r.name = genName()
	}

	names[r.name] = true
	return r.name, nil
}

func (r *Robot) Reset() {
	names[r.name] = false
	r.name = ""
}
