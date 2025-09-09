package bootstrap

import "github.com/wang900115/Opher/network"

var Network network.PeerManger

func InitializeNetWork(self *network.SelfPeer) {
	Network = network.NewGraph(self)
}
