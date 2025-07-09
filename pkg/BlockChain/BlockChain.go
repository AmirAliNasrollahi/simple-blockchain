package BlockChain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)


type Block struct {
	Timestamp       int64
	Trx             []Trx
	BeforeBlockHash string
	Nounce          string
}
type BlockChain struct {
	Blocks []Block
}


func InitBlockChain() BlockChain {
	BC := BlockChain{}
	return BC
}

/*
* the Blocks in Blockchain is immutable
* if the new block have created,
* you cant add new trx to before Blocks
* its amazing ...
 */
func (BC *BlockChain) MindBlock() {
	var block Block
	if len(BC.Blocks) == 0 {
		block = Block{
			Timestamp:       time.Now().Unix(),
			BeforeBlockHash: "",
		}
		BC.Blocks = append(BC.Blocks, block)
		return
	}
	// create a new Hasher variable
	hasher := sha256.New()

	// convert the last Block to json([]byte)
	byteOfLastBlock, _ := json.Marshal(BC.Blocks[len(BC.Blocks)-1])

	// pass the []byte to write a hash
	hasher.Write(byteOfLastBlock)

	block = Block{
		Timestamp: time.Now().Unix(),
		/*
		* in the sum method in the hasher
		* we have salt. when you pass the input
		* to the Sum method, its exactly salt
		 */
		BeforeBlockHash: string(hasher.Sum(nil)),
	}
	BC.Blocks = append(BC.Blocks, block)
}


func PutTrxFromMemPool() {

}



// the wallet address is a compressed of the public key. -> it mabey changed with some recursive algorithme
// the first,miners find the public key from wallet addr

// in the bitcoin it accept with the ECDSA algorithme
/**
the template of blockchain in bitcoin with real example

{
  "txid": "a1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d",
  "inputs": [
    {
      "address": "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
      "amount": 0.01005,
      "previous_txid": "xyz123..."  // تراکنش قبلی که این BTCها از آن آمده
    }
  ],
  "outputs": [
    {
      "address": "1P5ZEDWTKTFGxQjZphgWPQUpe554WKDfHQ",
      "amount": 0.01
    }
  ],
  "fee": 0.00005,
  "block_height": 824543
}


another example:
{
  "version": 1,
  "inputs": [
    {
      "prev_tx_hash": "abc123...", // UTXO اول
      "prev_out_index": 0,
      "scriptSig": "304502... (امضای مربوط به UTXO اول)",
      "sequence": 0xffffffff
    },
    {
      "prev_tx_hash": "def456...", // UTXO دوم
      "prev_out_index": 1,
      "scriptSig": "483045... (امضای مربوط به UTXO دوم)",
      "sequence": 0xffffffff
    }
  ],
  "outputs": [
    {
      "value": 0.05,
      "scriptPubKey": "76a914..." // آدرس گیرنده
    }
  ],
  "locktime": 0
}
**/
