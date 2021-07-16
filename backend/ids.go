package main

import (
	"fmt"
	"math/rand"
	"time"
)

var idRand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func newID(prefix string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"
	id := make([]byte, 18)
	for i := 0; i < len(id); i++ {
		id[i] = chars[idRand.Intn(len(chars))]
	}
	return fmt.Sprintf("%s%s", prefix, id)
}
