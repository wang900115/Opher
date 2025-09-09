package main

import "github.com/wang900115/Opher/bootstrap"

func main() {
	bootstrap.InitializeLevelDB()
	bootstrap.InitializeIPFS()
	// bootstrap.InitializeWallet()
	bootstrap.InitializeNetWork(nil)
	bootstrap.InitializeBlockChain()
}
