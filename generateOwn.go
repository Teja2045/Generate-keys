package main

import (
	"encoding/hex"
	"fmt"
	rsa "keys/userRSA"
	sha256 "keys/userSHA256"
	"math/big"
)

func PrivateKey() (*big.Int, *big.Int, *big.Int) {
	return rsa.PrivateKey()
}

func PublicKey(privateKey *big.Int, phi *big.Int) *big.Int {
	return rsa.PublicKey(privateKey, phi)
}

func Address(publicKey *big.Int) string {
	input := publicKey.String()                      // big integer to string
	paddedByteArrayInput := sha256.PaddedData(input) // string to bigInteger
	encodedString := hex.EncodeToString(sha256.Usersha256(paddedByteArrayInput))
	return encodedString // returns hexadecimal string as output
}

func main2() {
	privateKey, phi, _ := PrivateKey()
	fmt.Println(privateKey)
	publicKey := PublicKey(privateKey, phi)
	fmt.Println(publicKey)
	address := Address(publicKey)
	fmt.Println(address)

}
