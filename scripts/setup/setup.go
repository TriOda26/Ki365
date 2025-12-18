package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
)

type smap struct {
	k string
	v string
}

var s = []string{
	"KI365_KICAD_PASSWORD",
}

func main() {
	if fileExists(".env") {
		log.Println("WARNING: .env already exists! Only printing generated passwords:")
		generateRandomizedPasswords()
	} else {
		log.Println("Creating .env file with generated passwords")

		var m = []smap{}

		for _, val := range s {
			p, _ := generatePassword(20, true, true, true)
			m = append(m, smap{val, p})
		}

		file, err := os.OpenFile(".env", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		for i, val := range m {
			fmt.Printf("Pos %d, key %s, vak %s\n", i, val.k, val.v)
			s, b := fmt.Fprintf(file, "%s='%s'\n", val.k, val.v)
			if s == 0 {
				log.Println("WARNING: Printed zero length password, check .env for correct syntax!")
			}
			if b != nil {
				log.Fatal(err)
			}
		}

		log.Printf("Wrote .env file with %d generated passwords\n", len(m))
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if errors.Is(err, fs.ErrNotExist) {
		return false
	}

	return err == nil && !info.IsDir()
}

func generateRandomizedPasswords() error {
	for i := 0; i < 3; i++ {
		s, err := generatePassword(20, true, true, true)
		if err != nil {
			return err
		}
		log.Println(s)
	}
	return nil
}

func generatePassword(length int, useLetters bool, useSpecial bool, useNum bool) (string, error) {
	charset := ""
	var p []byte

	if useLetters {
		charset += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if useSpecial {
		// safe subset of special characters
		// charset += "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
		charset += "!@#%*()_+-=[]{}:,./?~"
	}
	if useNum {
		charset += "0123456789"
	}
	if charset == "" {
		return "", errors.New("must set one argument as true")
	}

	for i := 0; i < length; i++ {
		rand := rand.Intn(len(charset))
		p = append(p, charset[rand])
	}
	return string(p), nil
}
