package main

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

func main() {
	args := os.Args[1:]
    done := make(chan bool)

	var messages chan string = make(chan string)

	go Mine(args[0], args[1], messages, done)
	go Printer(messages)

    <-done
}

func Printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func Mine(input string, target string, messages chan string, done chan bool) {
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		nonce := r.Uint64()
		hash := GetHash(input + "." + S(nonce))
		message := fmt.Sprintf("%s\tHash: %s...", S(nonce), hash[0:32])
		messages <- message

		if CheckHash(hash, target) {
			fmt.Println("---------Found----------")
			fmt.Println(input + "." + S(nonce))
			fmt.Println(hash)

            os.Exit(0)
            done <- true
		}
	}
}

func CheckHash(hash string, target string) bool {
	if TrimHash(hash, utf8.RuneCountInString(target)) == target {
		return true
	}

	return false
}

func TrimHash(hash string, len int) string {
	trimmed := hash[0:len]

	return trimmed
}

func S(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func GetHash(s string) string {
	hasher := sha512.New()
	hasher.Write([]byte(s))
	result := hasher.Sum(nil)

	hash := fmt.Sprintf("%x", result)

	return hash
}
