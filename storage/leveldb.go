package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type LEVELDB interface {
	Put(key, value []byte) error    // Put stores a key-value pair.
	Get(key []byte) ([]byte, error) // Get retrieves the value by key.
	Delete(key []byte) error        // Delete removes a key-value pair.
	Exist(key []byte) (bool, error) // Checks if the key exist.
	Close() error                   // Cloase the db.
}

type LevelDB struct {
	db *leveldb.DB
}

func NewLevelDB(path string) LEVELDB {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic(err)
	}
	return &LevelDB{db: db}
}

func (l *LevelDB) Put(key, value []byte) error {
	return l.db.Put(key, value, nil)
}

func (l *LevelDB) Get(key []byte) ([]byte, error) {
	return l.db.Get(key, nil)
}

func (l *LevelDB) Delete(key []byte) error {
	return l.db.Delete(key, nil)
}

func (l *LevelDB) Exist(key []byte) (bool, error) {
	return l.db.Has(key, nil)
}

func (l *LevelDB) Close() error {
	return l.db.Close()
}
