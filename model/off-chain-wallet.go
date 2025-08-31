package model

import (
	"crypto/rsa"
)

// Wallet represents a user's cryptographic wallet.
type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}
