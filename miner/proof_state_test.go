package miner

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestProofStatePoolPut(t *testing.T) {
	//proofStatePool := NewProofStatePool()
	//
	//addrs := GenerateAddrs()
	//for i := 0; i < 11; i++ {
	//	//proofStatePool.Put(big.NewInt(1), addrs[0], addrs[i])
	//}
	//
	//for k, v := range proofStatePool.proofs {
	//	fmt.Println("current height:", k.String(), "===current count: ", v.count)
	//}
}

func TestOnlineValidator_GetAllAddress(t *testing.T) {
	var vals OnlineValidator
	vals = make(OnlineValidator)
	vals[common.HexToAddress("0x1000000000000000000000000000000000000000")] = struct{}{}
	vals[common.HexToAddress("0x1000000000000000000000000000000000000001")] = struct{}{}
	vals[common.HexToAddress("0x1000000000000000000000000000000000000002")] = struct{}{}
	vals[common.HexToAddress("0x1000000000000000000000000000000000000003")] = struct{}{}
	vals[common.HexToAddress("0x1000000000000000000000000000000000000004")] = struct{}{}
	addrs := vals.GetAllAddress()
	for _, addr := range addrs {
		t.Log(addr.Hex())
	}
}

func TestProofStatePool(t *testing.T) {
	proofStatePool := NewProofStatePool()

	height := big.NewInt(1)
	ps := newProofState(common.HexToAddress("0x1000000000000000000000000000000000000000"), common.HexToAddress("0x1000000000000000000000000000000000000000"))
	proofStatePool.proofs[height.Uint64()] = ps
	t.Log(proofStatePool.proofs[height.Uint64()] == nil)
	height2 := big.NewInt(1)
	ps2 := newProofState(common.HexToAddress("0x1000000000000000000000000000000000000001"), common.HexToAddress("0x1000000000000000000000000000000000000001"))
	proofStatePool.proofs[height2.Uint64()] = ps2
	t.Log(proofStatePool.proofs[height2.Uint64()] == nil)

	t.Log(proofStatePool.proofs[height2.Uint64()] == proofStatePool.proofs[height.Uint64()])

	//proofStatePool.ClearPrev(height)

	//for index := range proofStatePool.proofs {
	//	t.Log("index=", index)
	//}

}
