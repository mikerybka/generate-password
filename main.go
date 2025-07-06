package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	length := 12
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charset[rand.IntN(len(charset))]
	}
	fmt.Println(string(password))
}
