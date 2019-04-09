// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package node

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/IMTTHOLDINGCORP/go-aether/p2p"
	"github.com/IMTTHOLDINGCORP/go-aether/p2p/nat"
)

const (
	DefaultPort     = 60404
	DefaultHTTPHost = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort = 38545       // Default TCP port for the HTTP RPC server
	DefaultWSHost   = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort   = 38546       // Default TCP port for the websocket RPC server
)

// DefaultConfig contains reasonable default settings.
var DefaultConfig = Config{
	DataDir:     DefaultDataDir(),
	HTTPPort:    DefaultHTTPPort,
	HTTPModules: []string{"net", "web3"},
	WSPort:      DefaultWSPort,
	WSModules:   []string{"net", "web3"},
	P2P: p2p.Config{
		ListenAddr:      fmt.Sprintf(":%d", DefaultPort),
		DiscoveryV5Addr: ":44946",
		MaxPeers:        25,
		NAT:             nat.Any(),
	},
}

func DefaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "aether")
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "aether")
		} else {
			return filepath.Join(home, ".aether")
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

//add by liangc : for testnet build ipc path
func TestDataDir() string {
	home := homeDir()
	if home != "" {
		testnet := "testnet"
		if runtime.GOOS == "darwin" {
			return filepath.Join(home, "Library", "aether", testnet)
		} else if runtime.GOOS == "windows" {
			return filepath.Join(home, "AppData", "Roaming", "aether", testnet)
		} else {
			return filepath.Join(home, ".aether", testnet)
		}
	}
	return ""
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
