package main

import (
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println("Starting")
	args := os.Args[1:]
	target := args[1]
	inputHash := args[0]

	var i uint64

	for i = 0; i < 18446744073709551615; i++ {
		hash := GetHash(inputHash + "." + S(i))
		fmt.Println(hash)

		if CheckHash(hash, target) {
            fmt.Println("Found----------")
            fmt.Println(args[0] + "." + S(i))
			fmt.Println(hash)
			os.Exit(1)
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
	start := 128 - len
	trimmed := hash[start:128]
	//fmt.Println(trimmed)

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
