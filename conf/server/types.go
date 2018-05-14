package server

const (
	BlockChainMonitorUrl = "tcp://127.0.0.1:46757"
	Token                = "iris"
	InitConnectionNum    = 1000
	MaxConnectionNum     = 2000
	ChainId              = "pangu"
	SyncCron             = "0-59 * * * * *"

	SyncMaxGoroutine     = 2000
	SyncBlockNumFastSync = 2000
)
