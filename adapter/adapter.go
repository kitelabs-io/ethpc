package adapter

import (
	"github.com/kitelabs-io/ethrpc/adapter/ethereum"
)

func New(chainID uint, url string) (EthClientAdapter, error) {
	switch chainID {
	default:
		return ethereum.NewAdapter(url)
	}
}
