package main

import (
	"coin/blockchain"
	"fmt"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("test1")
	chain.AddBlock("test2")
	chain.AddBlock("test3")

	// chain_list := chain.AllBlock()
	for _, v := range chain.AllBlock() {
		fmt.Println(v)
		fmt.Println(v.Data)
		// result :
		// &{Genesis 81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5 }
		// Genesis
		// &{test1 c76474fc06647fc62e418fb5e11d1d512831ce0df85d09a303a6b4a162978e24 81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5}
		// test1
		// &{test2 04a942bbdbc0a477c14d32187592b6215a0ac4175090f5d4705ac92560a2fda9 c76474fc06647fc62e418fb5e11d1d512831ce0df85d09a303a6b4a162978e24}
		// test2
		// &{test3 0b891e744da8183cc9abc9465446236cd05309ed00f963bc541e600ed8518bcb 04a942bbdbc0a477c14d32187592b6215a0ac4175090f5d4705ac92560a2fda9}
		// test3

	}
}
