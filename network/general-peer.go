package network

// NetWork peer except self
type GeneralPeer struct {
	ID          string // Peer's Pubkey hash
	Address     string // Peer's IP:Port
	IsValidator bool   // Is validator for BTF
	LastSeen    int64  // HeartBreak timestamp (healthcheck)
	Height      int64  // Current local chain height
}
