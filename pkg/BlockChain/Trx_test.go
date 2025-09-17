package BlockChain

import "testing"

func findTrxByTrxIdTest(t testing.T) {
	firstInput := Input{
		Address: "TEST",	
		Amount: 10,
		Previous_trx: "k423opuiu423hgff234h2310",
	}
	secondInput := Input{
		Address: "newTEst",	
		Amount: 10,
		Previous_trx: "j4u3opuiuoiu342987h43kjl",
	}

	findTrxByTrxId(firstInput, secondInput)
	
	t.Log("the test passing")
}
