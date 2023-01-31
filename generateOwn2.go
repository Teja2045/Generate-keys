package main

import (
	"encoding/hex"
	"fmt"
	rsa "keys/userRSA2"
	sha256 "keys/userSHA256"
	"strconv"
)

func PrivateKey2(phi int, e int) int {
	return rsa.PrivateKey(phi, e)
}

func PublicKey2() (int, int) {
	phi, e := rsa.PublicKey()
	return phi, e
}

func Address2(publicKey int) string {
	input := strconv.Itoa(publicKey)                 // integer to string
	paddedByteArrayInput := sha256.PaddedData(input) // string to bigInteger
	encodedString := hex.EncodeToString(sha256.Usersha256(paddedByteArrayInput))
	return encodedString // returns hexadecimal string as output
}

func main() {
	phi, publicKey := PublicKey2()
	privateKey := PrivateKey2(phi, publicKey)
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	address := Address2(publicKey)
	fmt.Println(address)

}
