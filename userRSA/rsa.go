package userRSA

import (
	"math/big"
	"math/rand"
)

// GenerateKeys returns public and private keys

func PrivateKey() (*big.Int, *big.Int, *big.Int) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31,
		37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83,
		89, 97, 101, 103, 107, 109, 113, 127, 131, 137,
		139, 149, 151, 157, 163, 167, 173, 179, 181,
		191, 193, 197, 199, 211, 223, 227, 229, 233,
		239, 241, 251, 257, 263, 269, 271, 277, 281,
		283, 293, 307, 311, 313, 317, 331, 337, 347,
		349, 353, 359, 367, 373, 379, 383, 389, 397}
	randomIndex := rand.Intn(len(primes))
	randomPrimeA := primes[randomIndex]
	randomIndex = rand.Intn(len(primes))
	randomPrimeB := primes[randomIndex]
	p := big.NewInt(int64(randomPrimeA))
	q := big.NewInt(int64(randomPrimeB))
	n := new(big.Int).Mul(p, q)

	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	privateKey := big.NewInt(17)
	return privateKey, phi, n
}

func PublicKey(privateKey *big.Int, phi *big.Int) *big.Int {
	publicKey := new(big.Int).Exp(privateKey, big.NewInt(1), phi)
	return publicKey
}

func GenerateKeys() (*big.Int, *big.Int, *big.Int) {

	privateKey, phi, n := PrivateKey()

	publicKey := PublicKey(privateKey, phi)

	return publicKey, privateKey, n
}

// Encrypt encrypts a message
func Encrypt(message, publicKey, n *big.Int) *big.Int {
	cipher := new(big.Int).Exp(message, publicKey, n)
	return cipher
}

// Decrypt decrypts a message
func Decrypt(cipher, privateKey, n *big.Int) *big.Int {
	message := new(big.Int).Exp(cipher, privateKey, n)
	return message
}

// func main() {
// 	message := big.NewInt(42)
// 	publicKey, privateKey, n := GenerateKeys()

// 	cipher := Encrypt(message, publicKey, n)
// 	decryptedMessage := Decrypt(cipher, privateKey, n)

// 	fmt.Println("Message:", message)
// 	fmt.Println("Cipher: ", cipher)
// 	fmt.Println("Decrypted message:", decryptedMessage)
// }
