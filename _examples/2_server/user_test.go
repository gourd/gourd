package main

import (
	"math/rand"
	"testing"
)

func dummyNewUser() *User {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randSeq := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	return &User{
		Username: randSeq(10),
		Password: randSeq(8),
	}
}

func TestUser(t *testing.T) {
	var u OAuth2User = &User{}
	_ = u
}
