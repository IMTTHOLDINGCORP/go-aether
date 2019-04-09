// Copyright 2015 The go-ethereum Authors
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

package core

import "github.com/IMTTHOLDINGCORP/go-aether/common"

// BadHashes represent a set of manually tracked bad hashes (usually hard forks)
var BadHashes = map[common.Hash]bool{
	// add by liangc : testnet : rollback to 20 on testnet for test makecoin
	common.HexToHash("0xc6f3bc33763e07313a802f2f13c0596928e8a36cdc6ae33368f3ec5efa191b8e"): true,
	// add by liangc : testnet : rollback to 22 on testnet for test makecoin
	common.HexToHash("0x15e60cccd47c0ed664272f09420b4b0d6ce3fb8d2ea5f302e9c87ba960b990b6"): true,
	// add by liangc : testnet : rollback to 28 on testnet for test makecoin
	common.HexToHash("0xa78aec90360a3a279523de173f3a4bd7fdf3f0de86364b1d5d7787274af394c4"): true,
}
