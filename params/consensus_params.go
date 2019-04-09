package params

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/IMTTHOLDINGCORP/go-aether/common"
	"github.com/IMTTHOLDINGCORP/go-aether/log"
	"github.com/hashicorp/golang-lru"
	"math/big"
	"os"
	"sort"
)

type ChiefInfo struct {
	StartNumber *big.Int // 0 is nil
	Version     string
	Addr        common.Address
	Abi         string
}

func (self *ChiefInfo) String() string {
	return fmt.Sprintf("start: %d , vsn: %s , addr: %s", self.StartNumber.Int64(), self.Version, self.Addr.Hex())
}

type RegisterInfo struct {
	ChiefInfo
}

type RegisterInfoList []*RegisterInfo

func (p RegisterInfoList) Len() int { return len(p) }
func (p RegisterInfoList) Less(i, j int) bool {
	return p[i].StartNumber.Int64() > p[j].StartNumber.Int64()
}
func (p RegisterInfoList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ChiefInfoList []*ChiefInfo

func (p ChiefInfoList) Len() int { return len(p) }
func (p ChiefInfoList) Less(i, j int) bool {
	return p[i].StartNumber.Int64() > p[j].StartNumber.Int64()
}
func (p ChiefInfoList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func newRegisterInfo(num *big.Int, vsn string, addr common.Address, abi string) *RegisterInfo {
	r := &RegisterInfo{}
	r.StartNumber = num
	r.Version = vsn
	r.Addr = addr
	r.Abi = abi
	return r
}

func newChiefInfo(num *big.Int, vsn string, addr common.Address, abi string) *ChiefInfo {
	return &ChiefInfo{
		StartNumber: num,
		Version:     vsn,
		Addr:        addr,
		Abi:         abi,
	}
}

// for chief
// TribeChiefABI is the input ABI used to generate the binding from.
const (
	// copy from chieflib
	TribeChief_0_0_6ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"volunteer\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"volunteers\",\"type\":\"address[]\"}],\"name\":\"filterVolunteer\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"signerList\",\"type\":\"address[]\"},{\"name\":\"blackList\",\"type\":\"address[]\"},{\"name\":\"scoreList\",\"type\":\"uint256[]\"},{\"name\":\"numberList\",\"type\":\"uint256[]\"},{\"name\":\"totalVolunteer\",\"type\":\"uint256\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVolunteers\",\"outputs\":[{\"name\":\"volunteerList\",\"type\":\"address[]\"},{\"name\":\"weightList\",\"type\":\"uint256[]\"},{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignerLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fillVolunteerForTest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fillSignerForTest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"genesisSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

	Register_0_0_1ABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"Register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"h\",\"type\":\"string\"}],\"name\":\"appendBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"h\",\"type\":\"string\"}],\"name\":\"verifyBlackList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"h\",\"type\":\"string\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"append\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVsn\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getId\",\"outputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_cb\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"hasId\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"rawdata\",\"type\":\"string\"}],\"name\":\"appendId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	TribemapABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"delAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"existAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"makeCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"makeCoinLog\",\"outputs\":[{\"name\":\"numList\",\"type\":\"uint256[]\"},{\"name\":\"coinList\",\"type\":\"uint256[]\"},{\"name\":\"toList\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"addAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"
)

var (
	ChiefBaseBalance = new(big.Int).Mul(big.NewInt(1), big.NewInt(Finney))

	MboxChan            = make(chan Mbox, 32)
	TribemapService     = make(chan Mbox, 384)
	InitTribemapService = make(chan struct{})
	//close at tribe.init
	TribeReadyForAcceptTxs = make(chan struct{})
	InitTribe              = make(chan struct{})
	//close at tribeService
	InitTribeStatus                      = make(chan struct{})
	InitRegisterService                  = make(chan struct{})
	chiefInfoList       ChiefInfoList    = nil
	registerInfoList    RegisterInfoList = nil
	// added by cai.zhihong
	// ChiefTxGas = big.NewInt(400000)
	abiCache *lru.Cache = nil
)

func TribemapInfo(num *big.Int) (n *big.Int, addr common.Address) {
	if IsTestnet() {
		n = TestnetChainConfig.TribemapBlock
		if TestnetChainConfig.TribemapBlock.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.TribemapBlock.Cmp(num) <= 0 {
			addr = TestnetChainConfig.TribemapAddress
		}
	} else {
		n = MainnetChainConfig.TribemapBlock
		if MainnetChainConfig.TribemapBlock.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.TribemapBlock.Cmp(num) <= 0 {
			addr = MainnetChainConfig.TribemapAddress
		}
	}
	return
}

// if input num less then nr001block ,enable new role for chief.tx's gaspool
func IsNR001Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.NR001Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.NR001Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// new_rule_002 to change block period
// NR002Block must big then zero
func IsNR002Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.NR002Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.NR002Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.NR002Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.NR002Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

// add by liangc : 18-09-13 : incompatible HomesteadSigner begin at this number
func IsNR003Block(num *big.Int) bool {
	if IsTestnet() {
		if TestnetChainConfig.NR003Block.Cmp(big.NewInt(0)) > 0 && TestnetChainConfig.NR003Block.Cmp(num) <= 0 {
			return true
		}
	} else {
		if MainnetChainConfig.NR003Block.Cmp(big.NewInt(0)) > 0 && MainnetChainConfig.NR003Block.Cmp(num) <= 0 {
			return true
		}
	}
	return false
}

func registerAddressList() (list RegisterInfoList) {
	if registerInfoList != nil {
		return registerInfoList
	}
	if IsTestnet() {
		list = RegisterInfoList{
			newRegisterInfo(TestnetChainConfig.Register001Block, "0.0.1", TestnetChainConfig.Register001Address, Register_0_0_1ABI),
		}
	} else {
		list = RegisterInfoList{
			newRegisterInfo(MainnetChainConfig.Register001Block, "0.0.1", MainnetChainConfig.Register001Address, Register_0_0_1ABI),
		}
	}
	registerInfoList = list
	return
}

// startNumber and address must from chain's config
func chiefAddressList() (list ChiefInfoList) {
	if chiefInfoList != nil {
		return chiefInfoList
	}
	if IsTestnet() {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(TestnetChainConfig.Chief006Block, "0.0.6", TestnetChainConfig.Chief006Address, TribeChief_0_0_6ABI),
		}
	} else {
		list = ChiefInfoList{
			// at same account and block number to deploy this contract can be get the same address
			newChiefInfo(MainnetChainConfig.Chief006Block, "0.0.6", MainnetChainConfig.Chief006Address, TribeChief_0_0_6ABI),
		}
	}
	chiefInfoList = list
	return
}

func GetRegisterInfoByVsn(vsn string) *RegisterInfo {
	for _, ci := range registerAddressList() {
		if ci.StartNumber != nil && ci.StartNumber.Int64() > 0 && ci.Version == vsn {
			return ci
		}
	}
	return nil
}

func GetChiefInfoByVsn(vsn string) *ChiefInfo {
	for _, ci := range chiefAddressList() {
		if ci.StartNumber.Int64() > 0 && ci.Version == vsn {
			return ci
		}
	}
	return nil
}
func GetChiefInfo(blockNumber *big.Int) *ChiefInfo {
	if blockNumber != nil && blockNumber.Cmp(big.NewInt(0)) > 0 {
		return getChiefInfo(chiefAddressList(), blockNumber)
	}
	return nil
}
func getChiefInfo(list ChiefInfoList, blockNumber *big.Int) *ChiefInfo {
	// TODO sort once only
	sort.Sort(list)
	for _, c := range list {
		if blockNumber.Int64() >= c.StartNumber.Int64() {
			return c
		}
	}
	return nil
}

// skip verify difficulty on this block number
func IsChiefBlock(blockNumber *big.Int) bool {
	return isChiefBlock(chiefAddressList(), blockNumber)
}

func isChiefBlock(list ChiefInfoList, blockNumber *big.Int) bool {
	for _, ci := range list {
		//log.Info("isChief", "a", ci.StartNumber, "b", blockNumber)
		if ci.StartNumber.Cmp(blockNumber) == 0 {
			return true
		}
	}
	return false
}

func IsChiefUpdate(data []byte) bool {
	if len(data) < 4 {
		return false
	} else {
		if bytes.Equal(data[:4], []byte{28, 27, 135, 114}) {
			volunteer := common.Bytes2Hex(data[4:])
			if common.HexToAddress(volunteer) == common.HexToAddress("") {
				log.Debug("tribemap.ExistAddress.EmptyInput", "addr", common.HexToAddress(volunteer).Hex())
				return true
			} else {
				// TODO master append cache ,one data need check 3 timesz
				i, err := tribemapExitAddress(common.HexToAddress(volunteer))
				log.Debug("tribemap.ExistAddress", "addr", common.HexToAddress(volunteer).Hex(), "i", i, "err", err)
				if err != nil {
					switch err.Error() {
					case "skip":
						return true
					default:
						log.Error("IsChiefUpdate.tribemap.ExistAddress", "addr", common.HexToAddress(volunteer).Hex(), "err", err)
						return false
					}
				}
				if i > 0 {
					return true
				}
			}
		}
	}
	return false
}

func TribemapOwner() (common.Address, error) {
	<-InitTribemapService
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: "owner",
		Rtn:    rtn,
	}
	TribemapService <- m
	success := <-rtn
	if success.Success {
		return success.Entity.(common.Address), nil
	} else {
		return common.Address{}, success.Entity.(error)
	}
}

func tribemapExitAddress(addr common.Address) (int64, error) {
	select {
	case <-InitTribemapService:
		rtn := make(chan MBoxSuccess)
		m := Mbox{
			Method: "existAddress",
			Rtn:    rtn,
		}
		m.Params = map[string]interface{}{"addr": addr}
		TribemapService <- m
		success := <-rtn
		if success.Success {
			return success.Entity.(int64), nil
		} else {
			return 0, success.Entity.(error)
		}
	default:
		return 0, errors.New("skip")
	}
}

/*
func IsChiefUpdate2(data []byte) bool {
	if abiCache == nil {
		abiCache, _ = lru.New(10)
	}
	if len(data) < 4 {
		return false
	}
	dk := append(data[:4], []byte{0, 0, 0, 0}...)
	if abiCache.Contains(string(dk)) {
		return true
	}
	if len(data) > 4 {
		for _, ci := range chiefAddressList() {
			reader := strings.NewReader(ci.Abi)
			dec := json.NewDecoder(reader)
			var abi abi.ABI
			if err := dec.Decode(&abi); err != nil {
				panic(fmt.Errorf("chief_abi_error : vsn=%s", ci.Version))
			}
			buf, _ := abi.Pack("update", common.Address{})
			bk := append(buf[:4], []byte{0, 0, 0, 0}...)
			abiCache.Add(string(bk), string(bk))
			if bytes.Equal(data[0:4], buf[0:4]) {
				return true
			}
		}
	}
	return false
}
*/
func IsChiefAddress(addr common.Address) bool {
	return isChiefAddress(chiefAddressList(), addr)
}
func isChiefAddress(list ChiefInfoList, addr common.Address) bool {
	if addr == common.HexToAddress("0x") {
		log.Warn("--> isChiefAddress :: address_not_be_empty", "addr", addr)
		return false
	}
	for _, ci := range list {
		if ci.Addr == addr {
			return true
		}
	}
	return false
}

// chief service message box obj
type Mbox struct {
	Method string
	Rtn    chan MBoxSuccess
	Params map[string]interface{}
}

type MBoxSuccess struct {
	Success bool
	Entity  interface{}
}

// called by chief.GetStatus
func SendToMsgBoxWithHash(method string, hash common.Hash, number *big.Int) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	if number == nil || hash == common.HexToHash("0x") {
		panic(errors.New("hash and number can not nil"))
	}
	m.Params = map[string]interface{}{"hash": hash, "number": number}
	MboxChan <- m
	return rtn
}

func SendToMsgBoxForFilterVolunteer(hash common.Hash, number *big.Int, addr common.Address) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: "FilterVolunteer",
		Rtn:    rtn,
	}
	if number == nil || hash == common.HexToHash("0x") {
		panic(errors.New("hash and number can not nil"))
	}
	m.Params = map[string]interface{}{"hash": hash, "number": number, "address": addr}
	MboxChan <- m
	return rtn
}

// called by chief.Update
func SendToMsgBoxWithNumber(method string, number *big.Int) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	if number != nil {
		m.Params = map[string]interface{}{"number": number}
	}
	MboxChan <- m
	return rtn
}

func SendToMsgBox(method string) chan MBoxSuccess {
	rtn := make(chan MBoxSuccess)
	m := Mbox{
		Method: method,
		Rtn:    rtn,
	}
	MboxChan <- m
	return rtn
}

// clone from chief.getStatus return struct
// for return to tribe by channel
type ChiefStatus struct {
	VolunteerList  []common.Address
	SignerList     []common.Address
	BlackList      []common.Address // append by vsn 0.0.3
	ScoreList      []*big.Int
	NumberList     []*big.Int
	Number         *big.Int
	Epoch          *big.Int
	SignerLimit    *big.Int
	VolunteerLimit *big.Int
	TotalVolunteer *big.Int
}

// clone from chief.getVolunteers return struct
// for return to tribe by channel
// append by vsn 0.0.6
type ChiefVolunteers struct {
	VolunteerList []common.Address
	WeightList    []*big.Int
	Length        *big.Int
}

func GetIPCPath() string {
	return os.Getenv("IPCPATH")
}

func IsTestnet() bool {
	return os.Getenv("TESTNET") == "1"
}
