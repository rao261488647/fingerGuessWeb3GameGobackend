package config

import (
	"fingerGuessWeb3GameGoBackend/pkg/global"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(global.BlockConfig.RpcUrl)
	if err != nil {
		return nil, err
	}
	return client, nil
}
