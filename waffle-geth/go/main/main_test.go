package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"github.com/Ethworks/Waffle/simulator"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestName(t *testing.T) {
	sim, err := simulator.NewSimulator()
	if err != nil {
		log.Fatal(err)
	}

	deploy(sim)
	address := common.HexToAddress("Cca4391cF25d9773b4831E240b1Ddf9317B094cd")
	decodeString, err := hex.DecodeString("06fdde03")
	if err != nil {
		log.Fatal(err)
	}

	callMsg := ethereum.CallMsg{
		From:       common.Address{},
		To:         &address,
		Gas:        0,
		GasPrice:   nil,
		GasFeeCap:  nil,
		GasTipCap:  nil,
		Value:      nil,
		Data:       decodeString,
		AccessList: nil,
	}

	res, err := sim.Backend.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("dupa")
	fmt.Println(res)
}

func deploy(sim *simulator.Simulator) {
	tx1 := "f908cb8084342770c0830f42408080b9087960c0604052600d60808190526c2bb930b83832b21022ba3432b960991b60a090815261002e916000919061007a565b50604080518082019091526004808252630ae8aa8960e31b602090920191825261005a9160019161007a565b506002805460ff1916601217905534801561007457600080fd5b5061011b565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826100b057600085556100f6565b82601f106100c957805160ff19168380011785556100f6565b828001600101855582156100f6579182015b828111156100f65782518255916020019190600101906100db565b50610102929150610106565b5090565b5b808211156101025760008155600101610107565b61074f8061012a6000396000f3fe60806040526004361061009c5760003560e01c8063313ce56711610064578063313ce5671461020e57806370a082311461023957806395d89b411461026c578063a9059cbb14610281578063d0e30db0146102ba578063dd62ed3e146102c25761009c565b806306fdde03146100a1578063095ea7b31461012b57806318160ddd1461017857806323b872dd1461019f5780632e1a7d4d146101e2575b600080fd5b3480156100ad57600080fd5b506100b66102fd565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f05781810151838201526020016100d8565b50505050905090810190601f16801561011d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013757600080fd5b506101646004803603604081101561014e57600080fd5b506001600160a01b03813516906020013561038b565b604080519115158252519081900360200190f35b34801561018457600080fd5b5061018d6103f1565b60408051918252519081900360200190f35b3480156101ab57600080fd5b50610164600480360360608110156101c257600080fd5b506001600160a01b038135811691602081013590911690604001356103f5565b3480156101ee57600080fd5b5061020c6004803603602081101561020557600080fd5b503561056d565b005b34801561021a57600080fd5b50610223610624565b6040805160ff9092168252519081900360200190f35b34801561024557600080fd5b5061018d6004803603602081101561025c57600080fd5b50356001600160a01b031661062d565b34801561027857600080fd5b506100b661063f565b34801561028d57600080fd5b50610164600480360360408110156102a457600080fd5b506001600160a01b038135169060200135610699565b61020c6106ad565b3480156102ce57600080fd5b5061018d600480360360408110156102e557600080fd5b506001600160a01b03813581169160200135166106fc565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103835780601f1061035857610100808354040283529160200191610383565b820191906000526020600020905b81548152906001019060200180831161036657829003601f168201915b505050505081565b3360008181526004602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b4790565b6001600160a01b03831660009081526003602052604081205482111561043c576040805162461bcd60e51b8152602060048201526000602482015290519081900360640190fd5b6001600160a01b038416331480159061047a57506001600160a01b038416600090815260046020908152604080832033845290915290205460001914155b156104fc576001600160a01b03841660009081526004602090815260408083203384529091529020548211156104d1576040805162461bcd60e51b8152602060048201526000602482015290519081900360640190fd5b6001600160a01b03841660009081526004602090815260408083203384529091529020805483900390555b6001600160a01b03808516600081815260036020908152604080832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35060019392505050565b336000908152600360205260409020548111156105ab576040805162461bcd60e51b8152602060048201526000602482015290519081900360640190fd5b33600081815260036020526040808220805485900390555183156108fc0291849190818181858888f193505050501580156105ea573d6000803e3d6000fd5b5060408051828152905133917f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65919081900360200190a250565b60025460ff1681565b60036020526000908152604090205481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103835780601f1061035857610100808354040283529160200191610383565b60006106a63384846103f5565b9392505050565b33600081815260036020908152604091829020805434908101909155825190815291517fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9281900390910190a2565b60046020908152600092835260408084209091529082529020548156fea26469706673582212203cea2021c7d64d9a9474d6c0e2db3ce3e79ef65a1f84743117b8fccd6543a50064736f6c634300070500331ba0895b90bec98e76befd2b45573dee6925ef6eb049c7ff7e6a3718897113012e05a045658af94169ab259563f3f817e336d374ea2d863ff27722e0188ddb00bfeaf5"

	bytes, err := hex.DecodeString(tx1)
	if err != nil {
		log.Fatal(err)
	}

	tx := &types.Transaction{}
	err = tx.UnmarshalBinary(bytes)
	if err != nil {
		log.Fatal(err)
	}

	err = sim.Backend.SimulatedBackend.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	sim.Backend.Commit()

	receipt, err := sim.Backend.SimulatedBackend.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().String())
	fmt.Println(receipt.ContractAddress.String())
	fmt.Printf("status = %d\n", receipt.Status)
	fmt.Printf("gasused = %d\n", receipt.GasUsed)

}
