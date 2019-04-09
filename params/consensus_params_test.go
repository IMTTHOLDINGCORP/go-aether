package params

import (
	"fmt"
	"testing"
	"sort"
	"math/big"
	"github.com/IMTTHOLDINGCORP/go-aether/common"
)

var data = ChiefInfoList{
}

func TestChiefInfoList(t *testing.T) {
	t.Log(data)
	sort.Sort(data)
	t.Log(data)
}

func TestGetChiefAddress(t *testing.T) {
	sort.Sort(data)
	t.Log(getChiefInfo(data, big.NewInt(5)))
	t.Log(getChiefInfo(data, big.NewInt(50010)))
	t.Log(getChiefInfo(data, big.NewInt(20010)))
	t.Log(getChiefInfo(data, big.NewInt(30000)))
	t.Log(getChiefInfo(data, big.NewInt(40010)))
	t.Log(getChiefInfo(data, big.NewInt(80000)))
}

func TestIsChiefAddress(t *testing.T) {
	t.Log(isChiefAddress(data, common.HexToAddress("0x04")))
	t.Log(isChiefAddress(data, common.HexToAddress("0x0f")))
}

type ChiefStatus1 struct {
	NumberList []*big.Int
	BlackList  []*big.Int
}

type ChiefStatus2 struct {
	NumberList []*big.Int
	BlackList  []*big.Int
}

func TestFooBar(t *testing.T) {
	a := ChiefStatus1{[]*big.Int{big.NewInt(1)}, []*big.Int{big.NewInt(2)}}
	t.Log(a)
	b := ChiefStatus2(a)
	t.Log(b)
	var x []common.Address
	x = nil
	t.Log(x == nil)
}

func TestIsChiefUpdate(t *testing.T) {
	data := []byte{28, 27, 135, 114, 0, 0}
	t.Log(IsChiefUpdate(data))
	t.Log(IsChiefUpdate(data))
	t.Log(IsChiefUpdate(data))
	data = []byte{28, 27, 135, 115, 0, 0}
	t.Log(IsChiefUpdate(data))

}

func TestAddr(t *testing.T) {
	add1 := common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88aB49")
	add2 := common.HexToAddress("0xAd4c80164065a3c33dD2014908c7563eFf88Ab49")
	t.Log(add1 == add2)
}

func TestRABI(t *testing.T){
	fmt.Println(Register_0_0_1ABI)
}
