package BlockChain

type Input struct {
	Address      string
	Amount       float64
	Previous_trx string // to say this coin of where coming?
}

type Output struct {
	Address string
	Amount  float64
}

type Trx struct {
	Inputs  []Input
	Outputs []Output
	Fee     float64
}

func CheckTrxIsValid(trx Trx) {
	
}
