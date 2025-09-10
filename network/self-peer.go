package network

import (
	"sync"

	"github.com/wang900115/Opher/common"
)

type PeerRole int

const (
	RoleValidator PeerRole = iota
	RoleFullNode
	RoleLightNode
)

// Ownership network idendtifier
type SelfPeer struct {
	Address     common.NodeAddrAct // Self's IP:Port
	IsValidator bool               // Is validator for BTF
	PubKey      []byte             // pubKey
	PriKey      []byte             // priKey (store in local)
}

// Local network prototype
type Graph struct {
	mu        sync.RWMutex     // Avoiding race condition
	self      *SelfPeer        // Self node
	Neighbors map[string]*Peer // Dicovery neightbor node
}

func NewGraph(self *SelfPeer) PeerManger {
	return &Graph{
		self:      self,
		Neighbors: make(map[string]*Peer),
	}
}

func (g *Graph) Self() *SelfPeer {
	return g.self
}

// !todo error handling
func (g *Graph) Discovery(p *Peer) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Neighbors[p.ID] = p
}

func (g *Graph) ListPeers() []*Peer {
	g.mu.RLock()
	defer g.mu.RUnlock()
	result := make([]*Peer, 0, len(g.Neighbors))
	for _, p := range g.Neighbors {
		result = append(result, p)
	}
	return result
}
