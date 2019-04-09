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

package params

var MainnetBootnodes = []string{
	"enode://09128a9cd9d87c8d663cacaa7c05af33f21987ea010e0f3476a8dd27d8c0e31ff96ff7a1317f76f63325b5bcb4febbd8cf592186b28482e9b55a23e612b2f1e2@101.251.230.211:40401",
	"enode://55e4226bf4f11bb4d0da53171e5f49f0e8dc801cdb43d0a096d6c578cd2f103ff168b9df158eef3af98788a495fbd48bc1dff28dda89fd2cbf1f278af36f1cd0@101.251.230.213:40401",
}

var TestnetBootnodes = []string{
	"enode://8f93614fb0f21ee25426947fb1d5853ed4d85358370097662786514f92161600f3cb3ae66685ab0ff97c306ec76a4c48343c63e6968f32a08fe8c5cb015a136e@101.251.230.212:40401",
}

var RinkebyBootnodes = []string{}

var RinkebyV5Bootnodes = []string{}

var DiscoveryV5Bootnodes = []string{}
