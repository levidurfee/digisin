package main

import (
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
    "math/rand"
    "time"
)

func main() {
	fmt.Println("Starting")
	args := os.Args[1:]
	target := args[1]
	inputHash := args[0]

    Mine(inputHash, target)
}

func Mine(input string, target string) bool {
    var i uint64


	for i = 0; i < 18446744073709551615; i++ {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        nonce := r.Uint64()
        //fmt.Println(nonce)

		hash := GetHash(input + "." + S(nonce))
		fmt.Println(hash)

		if CheckHash(hash, target) {
            fmt.Println("---------Found----------")
            fmt.Println(input + "." + S(nonce))
			fmt.Println(hash)

            return true
		}
	}

    return false
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
