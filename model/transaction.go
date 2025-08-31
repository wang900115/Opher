package model

import "github.com/shopspring/decimal"

// TransactionOutput represents a single output in a transaction
type TransactionOutput struct {
	Index      int64             // The output's index within the transaction. Used by next transaction input to reference this output.
	Value      decimal.Decimal   // Amount assigned to the receiver.
	PubKey     []byte            // Receiver's public key.
	Spent      bool              // Boolean flag indicating whether this output has already been spent out.
	LockScript []byte            // Script to Lock this output(e.g., timelock, multi-signature).
	MetaData   map[string]string // Key-value data. e.g.,, token type, addition notes, or extra info.
}

//TransactionInput represents a single input in a transaction
type TransactionInput struct {
	PreTransactionID          []byte // ID of the previous transaction being spent.
	PreTransactionOutputIndex int64  // Index of the specifi ouput in the previous transaction.
	Signature                 []byte // Sender's signature proving ownership of the reference output.
	Nonce                     uint64 // Prevent replay attacks or enable Replace-By-Fee.
	Witness                   []byte // Multi-signature or smart contract verification.
	PubKey                    []byte // Sender's public key used to verify the signature.
}

// Transaction represents a record in a block body
type Transaction struct {
	ID        []byte              // Unique identifier of the transaction, usually a hash of the transaction's contents.
	Inputs    []TransactionInput  // List of TransactionInput, each referencing a previous unspent output.
	Outputs   []TransactionOutput // List of TransactionOutput, specifying new spendable outputs created by this transaction.
	Timestamp int64               // Unix timestamp when the transaction was created(used for ordering or timeLock validation).
}

// MemPool stores transactions waiting to be included in a block
type MemPool struct {
	PendingTransaction []*Transaction
}

// TransactionReceipt stores the result of a transaction
type TransactionReceipt struct {
	TransactionID []byte            // Hash of the transaction.
	Status        bool              // true = success, false = failed
	GasUsed       decimal.Decimal   // Amount of computational gas consumed.
	Logs          []ReceiptLog      // Event logs by contract.
	MetaData      map[string]string // Extra info(e.g., error messages, execution traces).
}

type ReceiptLog struct {
	Address []byte   // Contract address emitting the log.
	Topics  [][]byte // Indexed event params.
	Data    []byte   // Non-indexed event data.
}
