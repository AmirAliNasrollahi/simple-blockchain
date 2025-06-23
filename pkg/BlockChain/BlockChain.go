package BlockChain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Trx struct {
	Sender string
	Geter  string
	Amount float32
}
type Block struct {
	Timestamp       int64
	Trx             []Trx
	BeforeBlockHash string
}
type BlockChain struct {
	Blocks []Block
}

func InitBlockChain() BlockChain {
	BC := BlockChain{}	
	return BC
}

// the Blocks in Blockchain is immutable
// if the new block have created,
// you cant add new trx to before Blocks
// its amazing ...
func (BC *BlockChain) AddBlock() {
	var block Block
	if len(BC.Blocks) == 0 {
		block = Block{
			Timestamp: time.Now().Unix(),
			BeforeBlockHash: "",
		}
		BC.Blocks = append(BC.Blocks, block)
		return 
	}
	// create a new Hasher variable
	hasher := sha256.New()

	// convert the last Block to json([]byte) 
	byteOfLastBlock , _:= json.Marshal(BC.Blocks[len(BC.Blocks)-1])
	
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
