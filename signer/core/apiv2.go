// Copyright 2018 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	 "math/big"
)

type TypedData struct {
	Types 		map[string] interface{}		`json:"types"`
	PrimaryType	string						`json:"primaryType"`
	Domain 		map[string] interface{}		`json:"domain"`
	Message 	map[string]	interface{}		`json:"message"`
}

type EIP712Domain struct {
	Name 		string 			`json:"name"`
	Version		string 			`json:"version"`
	ChainId		big.Int			`json:"chainId"`
	Address		common.Address	`json:"address"`
	Salt		hexutil.Bytes	`json:"salt"`
}

type EncodedDataType struct {

}

type EncodedData struct {

}

// Typed data according to EIP712
//
// If the format "\x19\x01" ‖ domainSeparator ‖ hashStruct(message)` is not respected,
// an error is returned
func (api *SignerAPI) SignStructuredData(ctx context.Context, addr common.MixedcaseAddress, data TypedData) (hexutil.Bytes, error) {
	fmt.Println("what are you, data?", data)
	fmt.Println("types.EIP712Domain", data.Types["EIP712Domain"])
	fmt.Println("primaryType", data.PrimaryType)
	fmt.Println("message.from", data.Message["from"])
	return []byte("0xdeadbeef"), nil
}