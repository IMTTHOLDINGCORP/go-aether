package main

import (
	"encoding/hex"
	"fmt"
	"github.com/IMTTHOLDINGCORP/go-aether/crypto"
	"os"
)

func main() {
	//48e74753e6213f3fda91c5bd2023ab239553a1347c64a0368ec830ad317d215e
	fmt.Println(os.Args)
	prv, _ := crypto.GenerateKey()
	buf := crypto.FromECDSA(prv)
	if len(os.Args) > 1 {
		//data, _ := hex.DecodeString(os.Args[1])
		prv, _ = crypto.HexToECDSA(os.Args[1])
		buf = crypto.FromECDSA(prv)
	}

	s := hex.EncodeToString(buf)
	fmt.Println("privatekey:", s)
	fmt.Println("address:", crypto.PubkeyToAddress(prv.PublicKey).Hex())
}
