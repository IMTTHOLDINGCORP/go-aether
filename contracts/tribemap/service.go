package tribemap

import (
	"context"
	"errors"
	"github.com/IMTTHOLDINGCORP/go-aether/accounts/abi/bind"
	"github.com/IMTTHOLDINGCORP/go-aether/common"
	"github.com/IMTTHOLDINGCORP/go-aether/contracts/tribemap/lib"
	"github.com/IMTTHOLDINGCORP/go-aether/eth"
	"github.com/IMTTHOLDINGCORP/go-aether/ethclient"
	"github.com/IMTTHOLDINGCORP/go-aether/log"
	"github.com/IMTTHOLDINGCORP/go-aether/node"
	"github.com/IMTTHOLDINGCORP/go-aether/p2p"
	"github.com/IMTTHOLDINGCORP/go-aether/params"
	"github.com/IMTTHOLDINGCORP/go-aether/rpc"
	"math/big"
	"time"
)

type TribemapService struct {
	tribemap *tribemaplib.Tribemap
	ipcpath  string
	server   *p2p.Server // peers and nodekey ...
	quit     chan int
	client   *ethclient.Client
	ethereum *eth.Ethereum
}

var tribemapService *TribemapService

func NewTribemapService(ctx *node.ServiceContext) (node.Service, error) {
	var ethereum *eth.Ethereum
	ctx.Service(&ethereum)
	ipcpath := params.GetIPCPath()
	tribemapService = &TribemapService{
		quit:     make(chan int),
		ipcpath:  ipcpath,
		ethereum: ethereum,
	}
	go tribemapService.loop()
	return tribemapService, nil
}

func (self *TribemapService) Protocols() []p2p.Protocol { return nil }
func (self *TribemapService) APIs() []rpc.API           { return nil }
func (self *TribemapService) Start(server *p2p.Server) error {
	go func() {
		for {
			var (
				be = eth.NewContractBackend(self.ethereum.ApiBackend)
				cn = self.ethereum.BlockChain().CurrentBlock().Number()
			)
			mn, maddr := params.TribemapInfo(cn)
			if maddr != common.HexToAddress("") {
				contract, err := tribemaplib.NewTribemap(maddr, be)
				if err != nil {
					panic(err)
				}
				tribemapService.tribemap = contract
				close(params.InitTribemapService)
				log.Info("<<TribemapService.Start>> success ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			} else if cn.Cmp(mn) >= 0 {
				log.Info("<<TribemapService.Start>> cancel ", "cn", cn.Int64(), "tn", mn.Int64())
				return
			}
			log.Info("<<TribemapService.Start>> waiting... ", "cn", cn.Int64(), "tn", mn.Int64())
			<-time.After(3 * time.Second)
		}
	}()
	self.server = server
	return nil
}

func (self *TribemapService) Stop() error {
	close(self.quit)
	return nil
}

// ===============================================================================
// biz functions
// ===============================================================================

func GetTribemapService() (*TribemapService, error) {
	log.Debug("<<tribemap.service.GetTribemapService>>")
	select {
	case <-params.InitTribemapService:
		return tribemapService, nil
	default:
		return nil, errors.New("wait init")
	}
}

func (self *TribemapService) ExistAddress(addr common.Address) (*big.Int, error) {
	log.Debug("<<tribemap.service.ExistAddress>>")
	select {
	case <-params.InitTribemapService:
		var ctx = context.Background()
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		opts := new(bind.CallOptsWithNumber)
		opts.Context = ctx
		i, err := self.tribemap.ExistAddress(opts, addr)
		log.Debug("<<TribemapService.ExistAddress>>", "r", i, "err", err)
		return i, err
	default:
		return nil, errors.New("wait init")
	}
}

func (self *TribemapService) existAddress(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var addr common.Address
	addr = mbox.Params["addr"].(common.Address)
	log.Debug("mbox.params", "addr", addr.Hex())
	i, err := self.ExistAddress(addr)
	if err != nil {
		success.Success = false
		success.Entity = err
	} else {
		success.Entity = i.Int64()
	}
	mbox.Rtn <- success
}

func (self *TribemapService) owner(mbox params.Mbox) {
	success := params.MBoxSuccess{Success: true}
	var ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	opts := new(bind.CallOptsWithNumber)
	opts.Context = ctx
	o , err := self.tribemap.Owner(opts)
	if err != nil {
		success.Success = false
		success.Entity = err
	} else {
		success.Entity = o
	}
	mbox.Rtn <- success
}

func (self *TribemapService) loop() {
	for {
		select {
		case <-self.quit:
			break
		case mbox := <-params.TribemapService:
			switch mbox.Method {
			case "existAddress":
				self.existAddress(mbox)
			case "owner":
				self.owner(mbox)
			}
		}
	}
}
