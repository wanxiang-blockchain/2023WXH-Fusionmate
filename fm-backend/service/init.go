package service

import "github.com/FusionMate/fm-backend/common/log"

var loginAddresses map[string]loginAddress
var ChainClient *ChainAdaptor

func init() {
	var err error
	loginAddresses = make(map[string]loginAddress, 0)
	ChainClient, err = NewChainAdaptor()
	if err != nil {
		log.Fatal("[PANIC] create chain client error: %s")
	}
}
