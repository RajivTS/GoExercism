package diffiehellman

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	upperLimit := big.NewInt(0).Add(p, big.NewInt(-2))
	privateInt, err := rand.Int(rand.Reader, upperLimit)
	if err != nil {
		fmt.Printf("Error while getting random int: %s", err)
		return privateInt
	}
	return privateInt.Add(privateInt, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
