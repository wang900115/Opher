package model

import "github.com/shopspring/decimal"

// Token represents a fixed-supply blockchain token.
type Token struct {
	Name        string
	Symbol      string
	TotalSupply decimal.Decimal
	Balances    map[string]decimal.Decimal
}
