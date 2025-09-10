package common

type NetworkType string

const (
	NewEthernet NetworkType = "eth"
	NewWifi     NetworkType = "wifi"
	NewVPN      NetworkType = "vpn"
	NewSTUN     NetworkType = "stun"
	NewRelay    NetworkType = "relay"
)

type NodeAddr struct {
	IP       string
	Port     int
	Network  NetworkType
	IsPublic bool
	LastSeen int64
	Latency  int
}

type PeerRole int

const (
	RoleValidator PeerRole = iota
	RoleFullNode
	RoleLightNode
)

type SelfPeer struct {
	ID     string
	Addrs  []NodeAddr
	Role   PeerRole
	PubKey []byte
	PriKey []byte
}

type GeneralPeer struct {
	ID       string
	Addrs    []NodeAddr
	Role     PeerRole
	LastSeen int64
	Height   int64
	Latency  int
}

// type NodeAddrAct interface {
// 	Connect() error
// 	IsReachable() bool
// 	UpdateLatency(lat int)
// 	MarkSeen()
// }

// type Node struct {
// 	Addr NodeAddr
// }

// func (n *Node) Connect() error {
// 	// !TODO
// 	return nil
// }

// func (n *Node) IsReachable() bool {
// 	// !TODO
// 	return true
// }

// func (n *Node) UpdateLatency(lat int) {
// 	n.Addr.Latency = lat
// }

// func (n *Node) MarkSeen() {
// 	n.Addr.LastSeen = time.Now().Unix()
// }
