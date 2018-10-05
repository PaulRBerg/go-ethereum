#!/usr/bin/env bash

go install -v ./cmd/clef/
clef --networkid 4 --keystore /Users/paulrberg/Library/Ethereum/rinkeby/keystore --4bytedb /Users/paulrberg/go/src/github.com/ethereum/go-ethereum/cmd/clef/4byte.json --rules ./rules.js --rpc --advanced
