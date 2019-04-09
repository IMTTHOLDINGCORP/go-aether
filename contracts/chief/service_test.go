package chief

import (
	"context"
	"github.com/IMTTHOLDINGCORP/go-aether/common"
	"github.com/IMTTHOLDINGCORP/go-aether/ethclient"
	"math/big"
	"testing"
)

var (
	conn, _ = ethclient.Dial("/Users/liangc/Library/aether/testnet/aether.ipc")
	ctx     = context.Background()
)

func TestWriteTxData(t *testing.T) {

}

func TestReadTxData(t *testing.T) {
	b1, e := conn.BlockByNumber(ctx, big.NewInt(1))
	t.Log(1, e)
	t.Log(b1.Hash().Hex())

	tx, _, e := conn.TransactionByHash(ctx, common.HexToHash("0x2e8749e693029b0dfee120a6137c1c285b28dfa843e98cca23ad7e4ea522405d"))
	t.Log(2, e)
	t.Log(tx.Hash().Hex(), tx.Data())
}

func TestHello(t *testing.T) {
	c := make(chan int)
	close(c)
	t.Log(1)
	<-c
	t.Log(2)
	<-c
	t.Log(3)
	<-c
	t.Log(4)

}
