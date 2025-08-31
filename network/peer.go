package network

import "sync"

type PeerManger interface {
	Self() *SelfPeer
	Discovery(p *Peer)
	ListPeers() []*Peer
}

// NetWork peer except self
type Peer struct {
	ID          string // Peer's Pubkey hash
	Address     string // Peer's IP:Port
	IsValidator bool   // Is validator for BTF
	LastSeen    int64  // HeartBreak timestamp (healthcheck)
	Height      int64  // Current local chain height
}

// Ownership network idendtifier
type SelfPeer struct {
	ID          string // Self's PubKey hash
	Address     string // Self's IP:Port
	IsValidator bool   // Is validator for BTF
	PubKey      []byte // pubKey
	PriKey      []byte // priKey (store in local)
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
