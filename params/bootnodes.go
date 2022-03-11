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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://527a7280361271ab4ff5ce0fd36bfc15800a0f629b8f23fc9436905b1efa5d9976286d8a038d304df409febd42d56f4ca63d49f0054f2659469f69b5edead857@3.137.172.33:33101",
	"enode://9cf8c54550466ef4a30eae37559397e44a454256f09231170a8a253c62b4d83c6aa05ee6adf75554c1011b70f18a669aaa49847816c033956204fe72442a0336@3.16.78.14:33104",
	"enode://ba113e8ec9cf6c1dd399a0c5baf08570dd05f6e7f201e647bc40e6ed83e7e43b1d39cc2e5b02f20866e334dcfc6eba5518851e626bcac85d9ff89dd80c6a5776@3.144.106.111:33107",
	"enode://7d12191a73d1b56f7102a7497683d2dc67b4b272d4609fe7d8bfb2b3b2f402a02952dd97834c9457f65a7e3abae2b5b38982c1bb4feef3587cae2a4d4395077e@3.14.129.214:33111",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the testnet.
var TestnetBootnodes = []string{
	"enode://10638ced559d8f4fa7e569097c312b930ea95d7be6d99d74a7ddfce7c141ae3c7d55d8485581d4260dedf96325e4ac89a904ff9928996c60fb0270c309233f03@3.144.195.4:33601",
	"enode://22f76f4f33373714cf6cb699fd1fcb1fe62973cf135479f887266ca704fc5b66525efa854a89bd49916ca010bfb363432032b465d859bce9760cbb8d9375bc5f@3.144.195.4:33602",
	"enode://c7aba15dfcf7d5de2939f802b400ec3ea5763d642dc167af3d7df5771e1a92850c2cac533ba06f66344abf56572ef8bb04403a110d23a3c4c14afd5008d2bb3f@3.144.195.4:33603",
}

var V5Bootnodes = []string{}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
