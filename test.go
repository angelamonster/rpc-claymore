package main

import (
	"fmt"
	//"net/rpc/jsonrpc"
	//"strconv"
	//"strings"
	claymore "./"
)

func main() {
	miner := claymore.Miner{Address: "w0004:3333"}
	info, err := miner.GetInfo()

	if err != nil {
		//log.Fatal(err)
		fmt.Printf(err)
	}
	fmt.Printf(info.MainCrypto.HashRate)
}
