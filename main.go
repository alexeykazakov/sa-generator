package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type SA struct {
	Name    string   `json:"name"`
	ID      string   `json:"id"`
	Secrets []string `json:"secrets"`
}

type SaSecret struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	saName := flag.String("name", "fabric8-gemini-server", "Analytics gemini server")
	flag.Parse()

	fmt.Println("\nPreview:")
	generateSA(*saName)
	fmt.Println("\nProd:")
	generateSA(*saName)
	fmt.Println("\nLocal:")
	generateSA(*saName)
}

func generateSA(saName string) {
	secret := RandStringRunes(50)
	id := uuid.NewV4().String()
	saSecret := SaSecret{
		Name:   saName,
		ID:     id,
		Secret: secret,
	}
	b, err := json.MarshalIndent(saSecret, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", string(b))

	h := hash(secret)
	sa := SA{
		Name:    saName,
		ID:      id,
		Secrets: []string{h},
	}

	b, err = json.MarshalIndent(sa, "        ", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n        %s\n", string(b))
}

func hash(password string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(h)
}
