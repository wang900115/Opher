package bootstrap

import "github.com/wang900115/Opher/storage"

var DB storage.LEVELDB
var IPFSNode storage.IPFS

// store on-chain info (e.g., block, state ...)
func InitializeLevelDB() {
	DB = storage.NewLevelDB("/db")
}

// store off-chain info(e.g., big files ...)
func InitializeIPFS() {
	IPFSNode = storage.NewIPFSNode("localhost:5001")
}
