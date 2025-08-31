package model

import (
	"github.com/shopspring/decimal"
)

// Account represents a blockchain account (either an EOA or a Contract Account).
type Account struct {
	Address []byte
	Nonce   uint64
	Balance decimal.Decimal
}

// ContractAccount represnets a smart contract account.
type ContractAccount struct {
	Address     []byte // creatorAddress + creatorNonce hash
	StorageRoot []byte // smart contract state's merkle root
	CodeHash    []byte // smart contract's hash
	Code        []byte // smart contract code
}
