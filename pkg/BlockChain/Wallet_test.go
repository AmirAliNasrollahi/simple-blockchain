package BlockChain

import "testing"


func TestWalletCreation(t *testing.T) {
	// arrange
	userName := "reza"
	// act
	output := CreateWallet(userName)	
	// assert
	t.Log(output)	
}
