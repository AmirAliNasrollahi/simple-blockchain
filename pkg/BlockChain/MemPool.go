package BlockChain

type MemPool struct {
	Trx []Trx
}

/*
* every Trx should first go to the Mempool, then
* the miner check which one is valid and then
* put out the valid transAction while the size of
* block going to be smaller than 1MB
*/

func PutTrxToMemPool() {
}

