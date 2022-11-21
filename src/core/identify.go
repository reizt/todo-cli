package core

import (
	"crypto/rand"
	"fmt"
)

func newId() string {
	idLen := 3
	b := make([]byte, idLen)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}
