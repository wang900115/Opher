package storage

import (
	"bytes"
	"io"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFS interface {
	Put(data []byte) (hash []byte, err error)          // Put data and return the generated hash
	Get(hash []byte) (data []byte, err error)          // Get data by hash
	Pin(hash []byte) error                             // Pin a hash to persist it locally
	UnPin(hash []byte) error                           // Unpin a hash to remove local persistence
	IsPinned(hash []byte) (bool, error)                // Check if a hash is pinned locally
	ListPinned() ([][]byte, error)                     // List all pinned hashes
	AddFile(reader io.Reader) (hash []byte, err error) // Add a file from a reader
	GetFile(hash []byte) (io.Reader, error)            // Get a file as a reader
	Remove(hash []byte) error                          // Remove a hash from local Node
	Size(hash []byte) (int64, error)                   // Get size of object by hash
	ResolveName(name string) ([]byte, error)           // Resolve a human-readable name to a hash(IPNS)
}

type IPFSNode struct {
	s *shell.Shell
}

func NewIPFSNode(url string) IPFS {
	return &IPFSNode{
		s: shell.NewShell(url),
	}
}

func (node *IPFSNode) Put(data []byte) (hash []byte, err error) {
	hashStr, err := node.s.Add(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return []byte(hashStr), nil
}

func (node *IPFSNode) Get(hash []byte) (data []byte, err error) {
	reader, err := node.s.Cat(string(hash))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (node *IPFSNode) Pin(hash []byte) error {
	return node.s.Pin(string(hash))
}

func (node *IPFSNode) UnPin(hash []byte) error {
	return node.s.Unpin(string(hash))
}

func (node *IPFSNode) IsPinned(hash []byte) (bool, error) {
	pins, err := node.s.Pins()
	if err != nil {
		return false, err
	}
	_, ok := pins[string(hash)]
	return ok, nil
}

func (node *IPFSNode) ListPinned() ([][]byte, error) {
	pinned, err := node.s.Pins()
	if err != nil {
		return nil, err
	}
	result := make([][]byte, 0, len(pinned))
	for h := range pinned {
		result = append(result, []byte(h))
	}
	return result, nil
}

func (node *IPFSNode) AddFile(reader io.Reader) (hash []byte, err error) {
	hashStr, err := node.s.Add(reader, nil)
	if err != nil {
		return nil, err
	}
	return []byte(hashStr), nil
}

func (node *IPFSNode) GetFile(hash []byte) (io.Reader, error) {
	reader, err := node.s.Cat(string(hash))
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func (node *IPFSNode) Remove(hash []byte) error {
	if err := node.s.Unpin(string(hash)); err != nil {
		return err
	}
	return nil
}

func (node *IPFSNode) Size(hash []byte) (int64, error) {
	stat, err := node.s.ObjectStat(string(hash))
	if err != nil {
		return 0, err
	}
	return int64(stat.CumulativeSize), nil
}

func (node *IPFSNode) ResolveName(name string) ([]byte, error) {
	hashStr, err := node.s.Resolve(name)
	if err != nil {
		return nil, err
	}
	return []byte(hashStr), nil
}
