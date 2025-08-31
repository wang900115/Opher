package model

import (
	"github.com/wang900115/Opher/network"
	"github.com/wang900115/Opher/storage"
)

// Initial block chain
type Chain struct {
	LastHash   []byte             // Hash of the latest block.
	Blocks     map[string]*Block  // In-Memory index for quick lookup.
	DBLEVELDB  storage.LEVELDB    // Persisent stroage on-chain interface.
	DBIPFS     storage.IPFS       // Persisent storage off-chain interface.
	Peers      network.PeerManger // NetWork layer interface.
	Difficulty int64              // Current mining difficulty.
	Validators [][]byte           // BFT validator's public keys.
	Round      int64              // Current BFT round,
	Epoch      int64              // Epoch for validator set,
	Proposer   []byte             // Current block proposer,
	StateRoot  []byte             // Root hash of world state,
	Height     int64              // Current blockchain height.
	Genesis    []byte             // Gensis block,
	Mempool    *MemPool           // Pending transactions,
}
