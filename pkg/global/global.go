package global

import (
	"fingerGuessWeb3GameGoBackend/config/setting"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

var (
	DbConfig    *setting.DataBase
	BlockConfig *setting.BlockConfig
	DbEngine    *gorm.DB
	RpcClient   *ethclient.Client
)
