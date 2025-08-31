package bootstrap

import "github.com/wang900115/Opher/model"

func InitializeBlockChain() {
	memPool := &model.MemPool{PendingTransaction: []*model.Transaction{}}

	Chain := &model.Chain{
		LastHash:   nil,
		Blocks:     make(map[string]*model.Block),
		DBLEVELDB:  DB,
		DBIPFS:     IPFSNode,
		Difficulty: 1,
		Validators: [][]byte{},
		Round:      0,
		Epoch:      0,
		StateRoot:  nil,
		Height:     0,
		Mempool:    memPool,
	}

	genesis := &model.Block{
		Header: model.BaseInfo{
			Version:   1,
			Height:    0,
			PreHash:   nil,
			Timestamp: 0,
		},
		Body: model.Payload{
			MetaData: map[string]string{"gensis": "true"},
		},
	}

	Chain.LastHash = genesis.Header.Hash
	Chain.Height = 1
}
