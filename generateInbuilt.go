package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func PrivateKey1() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func PublicKey1(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	publicKey := privateKey.PublicKey
	return &publicKey
}

func Adress1(publicKey *rsa.PublicKey) string {
	hash := sha256.Sum256([]byte(publicKey.N.String()))

	address := hex.EncodeToString(hash[:])
	return address
}

func main1() {
	privateKey := PrivateKey1()
	fmt.Println(privateKey)
	fmt.Println()
	publicKey := PublicKey1(privateKey)
	fmt.Println(publicKey)
	fmt.Println()
	address := Adress1(publicKey)
	fmt.Println(address)

}
