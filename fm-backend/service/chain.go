package service

import (
	"fmt"
	"github.com/FusionMate/fm-backend/abi/AssistantFactory"
	"github.com/FusionMate/fm-backend/conf"
	"github.com/FusionMate/fm-backend/dao"
	"github.com/FusionMate/fm-backend/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type ChainAdaptor struct {
	Client  *ethclient.Client
	Factory *AssistantFactory.AssistantFactory
}

func NewChainAdaptor() (*ChainAdaptor, error) {
	rpcUrl := conf.GConfig.GetString("contract.ethNode")
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	factory, err := AssistantFactory.NewAssistantFactory(
		common.HexToAddress(conf.GConfig.GetString("contract.assistantFactoryContractAddress")),
		client)
	if err != nil {
		return nil, err
	}
	return &ChainAdaptor{
		Client:  client,
		Factory: factory,
	}, nil
}

func (c *ChainAdaptor) GetNFTContractAddress(collectionId int64) (string, error) {
	address, err := c.Factory.AssistantFactoryCaller.AssistantsMap(
		&bind.CallOpts{
			Pending:     false,
			From:        common.Address{},
			BlockNumber: nil,
			Context:     nil,
		}, big.NewInt(collectionId))
	if err != nil {
		return "", err
	}
	return address.String(), nil
}

func GetAndUptNFTContractAddress(collectionID int64) error {
	addr, err := ChainClient.GetNFTContractAddress(collectionID)
	if err != nil {
		return err
	}
	if len(addr) == 0 {
		return fmt.Errorf("get nft contract address is empty")
	}
	err = dao.UptCollectionAddress(&model.Collection{
		CollectionID: collectionID,
		ContractAddr: addr,
	})
	if err != nil {
		return err
	}
	return nil
}
