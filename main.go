package main

import (
	"crypto/sha512"
	"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

var i int = 0
var start = time.Now()

func main() {
	args := os.Args[1:]
	done := make(chan bool)
	messages := make(chan string)

    tm.Clear()

	go Mine(args[0], args[1], messages, done)
	go Printer(messages)

	<-done
}

func Printer(c chan string) {
    x := 1
	for {
        if x > 10 {
            x = 1
        }
        tm.MoveCursor(1, x)

		msg := <-c
		tm.Println(msg)
        tm.Flush()
        x++
	}
}

func Mine(input string, target string, messages chan string, done chan bool) {
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		nonce := r.Uint64()
		hash := GetHash(input + "." + S(nonce))
		hps := getHashesPerSecond()
		message := fmt.Sprintf("%s\tHash: %s... %d/hps", S(nonce), hash[0:32], hps)
		messages <- message

		if CheckHash(hash, target) {
			t := time.Now()
			elapsed := t.Sub(start)

            tm.Clear()
            tm.MoveCursor(1, 1)

			tm.Println("---------Found----------")
			tm.Println("Attempts:\t" + strconv.Itoa(i))
			tm.Println("Duration:\t" + elapsed.String())
			tm.Println()
			tm.Println("Input:\t\t" + input + "." + S(nonce))
			tm.Println("Hash:\t\t" + hash)
            tm.Flush()

            done <- true

			os.Exit(0)
		}
	}
}

func getHashesPerSecond() int {
	i++
	t := time.Now()
	elapsed := t.Sub(start)
	elapsedSeconds := int(elapsed.Seconds())
	if elapsedSeconds > 0 {
		hps := i / elapsedSeconds
		return hps
	}

	return 0
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
