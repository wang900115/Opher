package model

import "math/big"

// Block basic structure
type Block struct {
	Header     BaseInfo        // Basic block information.
	Body       Payload         // Transactions and optional metadata.
	Validation HybridConsensus // Consensus and security-realated info.
}

// Describe the fundamental linking information of the block
type BaseInfo struct {
	Version                int64  // Block version (for upgrades and protocal compatibility).
	Height                 int64  // Block height (index in the chain, starting from genesis).
	Hash                   []byte // This block hash.
	PreHash                []byte // Hash of the prevoius block.
	MerkleStateRoot        []byte // Root hash of the world state.
	MerkleTransactionsRoot []byte // Root hash of all transactions.
	MerkleReceiptsRoot     []byte // Root hash of all transactions result.
	Timestamp              int64  // Unix timestamp when the block was created.
}

// Contains the actual block data
type Payload struct {
	MetaData     map[string]string // Key-value data. e.g., tags, or extra info.
	Transactions []Transaction     // List of transactions.
}

// ProofOfWork represents the PoW part of consensus
type ProofOfWork struct {
	Nonce  int64
	Target *big.Int // Target difficulty threshold
}

const (
	VoteNil = iota
	VotePrecommit
	VoteCommit
	VoteReject
)

// ByzantineFaultTolerance represents the BFT part of consensus
type ByzantineFaultTolerance struct {
	ValidatorID []byte // Validator's public key
	Signature   []byte // Validator's digital signature
	Vote        int    // Voting type (0= nil, 1=Precommit, 2=Commit, 3=Reject)
	Round       int64  // Current BFT Round
	Timestamp   int64  // Time of Vote
}

// HybridConsensus combines PoW + BFT
type HybridConsensus struct {
	Proposer   []byte                    // Miner/validator who proposed the block.
	Difficulty int64                     // Difficulty at the time of mining.
	Size       int64                     // Block size in bytes.
	PoW        ProofOfWork               // PoW info.
	BFT        []ByzantineFaultTolerance // BFT validators' votes
	State      string                    // e.g. "Pending", "Finalized", "Rejected"
}

// BlockPool stores blocks waiting to be included in a chain
type BlockPool struct {
	PendingBlocks map[string]*Block // stroes blocks not yet on-chain
}
