// Copyright 2017 The go-ethereum Authors
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

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const testnetAllocData = "\xe3\xe2\x94O\xc8n*\xbb\t\xbaR\x94M\x04\x0f\xd0~ \x06\x15\xa2\x00\x80\x8c O\xce^>%\x02a\x10\x00\x00\x00"
const mainnetAllocData = "\xe0\u07d46\xc9\xfb\xf1\xab\x10\x83\u06af\x86\xdaV2\\\x94\xa5\x89\xe2\xb0J\x89g\xb45D}\xdaD\x00\x00"
const rinkebyAllocData = ""
