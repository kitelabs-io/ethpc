package ethereum

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	adaptertypes "github.com/kitelabs-io/ethrpc/adapter/types"
)

func TestAdapterTestSuite(t *testing.T) {
	suite.Run(t, new(AdapterTestSuite))
}

type AdapterTestSuite struct {
	suite.Suite

	adapter *Adapter
}

func (ts *AdapterTestSuite) SetupTest() {
	adapter, err := NewAdapter("wss://ethereum.publicnode.com")

	assert.Nil(ts.T(), err)

	ts.adapter = adapter
}

func (ts *AdapterTestSuite) TestSubscribeNewHead() {
	ctx := context.Background()
	headerChannel := make(chan *adaptertypes.Header)

	sub, err := ts.adapter.SubscribeNewHead(ctx, headerChannel)
	defer sub.Unsubscribe()

	assert.Nil(ts.T(), err)

	go func() {
		defer close(headerChannel)

		for {
			select {
			case <-ctx.Done():
				return
			case err = <-sub.Err():
				assert.Error(ts.T(), err)

				return
			case originHeader := <-headerChannel:
				fmt.Printf("received header: %s\n", originHeader.Hash.String())
				assert.NotNil(ts.T(), originHeader)
			}
		}
	}()

	time.Sleep(12 * time.Second)
}
