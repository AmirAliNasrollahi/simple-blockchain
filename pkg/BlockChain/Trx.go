package BlockChain

// we save someCookie in the clientSide, something like
// Previous_trx
type Input struct {
	Previous_trx string // to say this coin of where coming?
	scriptSig    string // trx sign
}

type Output struct {
	Address string
	Amount  float64
}

type Trx struct {
	trx_id  string
	Inputs  []Input
	Outputs []Output
	Fee     float64
}

func CheckTrxIsValid(trx Trx) {

}

// why we cant use this form ?
// because this function create Trx, so we cant pass Trx
// to this and do calculate on it
//
//	func (trx *Trx) CreateTrx(input Input, output Output) {
//		trx.Fee = input.Amount - output.Amount
//	}
func CreateTrx(inputs []Input, outputs []Output, wallet string) {
	newTrx := &Trx{}
	// find before block which the Coin comming to sender
	newTrx.Inputs = inputs
	newTrx.Outputs = outputs

	// check if it has more than 1 inputs, it want to combinition the UTXO
	if len(inputs) > 1 {
		// for _, _= range inputs {
		// check the UTXO's valid or not
		// }
	}
}

func coinCombinition(senderWallet string, geterWallet string) {
	if senderWallet == geterWallet {

	}
}

func findTrxByTrxId(previous_trx string) (string, float64) {
	// it just for e.g
	wallet, amount := "8493j5m32k534klkj", 12.43
	return wallet, amount
}

func feeCalculation(sendingCoin float64) float64 {
	var fee float64
	if sendingCoin > 10 {
		fee = sendingCoin * 0.01
		return fee
	}
	fee = 0.1
	return fee

}
