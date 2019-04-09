package register

import (
	"context"
	"errors"
	"github.com/IMTTHOLDINGCORP/go-aether/accounts/abi/bind"
	"github.com/IMTTHOLDINGCORP/go-aether/common"
	"github.com/IMTTHOLDINGCORP/go-aether/contracts/register/lib"
	"github.com/IMTTHOLDINGCORP/go-aether/eth"
	"github.com/IMTTHOLDINGCORP/go-aether/ethclient"
	"github.com/IMTTHOLDINGCORP/go-aether/internal/ethapi"
	"github.com/IMTTHOLDINGCORP/go-aether/log"
	"github.com/IMTTHOLDINGCORP/go-aether/node"
	"github.com/IMTTHOLDINGCORP/go-aether/p2p"
	"github.com/IMTTHOLDINGCORP/go-aether/params"
	"github.com/IMTTHOLDINGCORP/go-aether/rpc"
	"time"
)

type RegisterService struct {
	register_0_0_1 *registerlib.Register_0_0_1
	ipcpath        string
	server         *p2p.Server // peers and nodekey ...
	quit           chan int
	client         *ethclient.Client
	ethereum       *eth.Ethereum
}

var registerService *RegisterService

func NewRegisterService(ctx *node.ServiceContext) (node.Service, error) {
	var apiBackend ethapi.Backend
	var ethereum *eth.Ethereum
	if err := ctx.Service(&ethereum); err == nil {
		apiBackend = ethereum.ApiBackend
	} else {
		return nil, errors.New("not support light node")
	}
	ipcpath := params.GetIPCPath()
	registerService = &RegisterService{
		quit:     make(chan int),
		ipcpath:  ipcpath,
		ethereum: ethereum,
	}
	if v0_0_1 := params.GetRegisterInfoByVsn("0.0.1"); v0_0_1 != nil {
		contract_0_0_1, err := registerlib.NewRegister_0_0_1(v0_0_1.Addr, eth.NewContractBackend(apiBackend))
		if err != nil {
			return nil, err
		}
		registerService.register_0_0_1 = contract_0_0_1
	}
	return registerService, nil
}

func (self *RegisterService) Protocols() []p2p.Protocol { return nil }
func (self *RegisterService) APIs() []rpc.API           { return nil }
func (self *RegisterService) Start(server *p2p.Server) error {
	self.server = server
	close(params.InitRegisterService)
	return nil
}

func (self *RegisterService) Stop() error {
	close(self.quit)
	return nil
}

// ===============================================================================
// biz functions
// ===============================================================================

func GetRegisterService() (*RegisterService, error) {
	select {
	case <-params.InitRegisterService:
		return registerService, nil
	default:
		return nil, errors.New("wait init")
	}
}

func (self *RegisterService) GetId(id string) (common.Address, error) {
	var (
		idVo struct {
			Id string
			Cb common.Address
		}
		ctx = context.Background()
	)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	opts := new(bind.CallOptsWithNumber)
	opts.Context = ctx
	idVo, err := self.register_0_0_1.GetId(opts, id)
	if err != nil {
		return common.Address{}, err
	}
	return idVo.Cb, nil
}

func (self *RegisterService) HasId(id string) (bool, error) {
	var (
		has bool
		ctx = context.Background()
	)
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	opts := new(bind.CallOptsWithNumber)
	opts.Context = ctx
	has, err := self.register_0_0_1.HasId(opts, id)
	if err != nil {
		return false, err
	}
	return has, nil
}

func (self *RegisterService) initEthclient() error {
	if self.client == nil {
		ethclient, err := ethclient.Dial(self.ipcpath)
		if err != nil {
			log.Error("ipc error ", "err", err)
			return err
		}
		self.client = ethclient
	}
	return nil
}
