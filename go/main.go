package main

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

const genesisMultihashString = "1220107d821c25dc0735200249df94a8bebc9c8e489744f86a4ca8919e81f19dcd72"

func main() {
	genesisMh, err := multihash.FromHexString(genesisMultihashString)
	if err != nil {
		panic(err)
	}
	genesisCid := cid.NewCidV1(cid.DagCBOR, genesisMh)
	fmt.Println("Genesis CID:", genesisCid)
}
