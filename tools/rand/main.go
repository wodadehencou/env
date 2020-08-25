package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func main() {
	n := flag.Int("n", 0, "number of random bytes to generate")
	s := flag.Int("s", 0, "length of random string to generate")
	flag.Parse()

	if *n != 0 {
		data := make([]byte, *n)
		rand.Read(data)

		for i, d := range data {
			fmt.Printf("0x%02X, ", d)
			if i%8 == 7 {
				fmt.Println()
			}
		}
		fmt.Println()
	}

	if *s != 0 {
		var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		b := make([]rune, *s)
		for i := range b {
			r, _ := rand.Int(rand.Reader, new(big.Int).SetUint64(uint64(len(letterRunes))))
			b[i] = letterRunes[r.Int64()]
		}
		fmt.Println(string(b))
	}
}
