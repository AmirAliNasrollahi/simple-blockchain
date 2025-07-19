package Handlers

import (
	"crypto_currency/pkg/BlockChain"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SignUp struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}
type Response struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world Bro")
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	user := &SignUp{}
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("\n we have an error when reading input\n")
		return
	}

	err = json.Unmarshal(body, user)

	if err != nil {
		log.Printf("we have an error when decoding the input \n")
		return
	}

	privateKeyBlock, publicKeyBlock := BlockChain.CreateWallet(user.Name + user.LastName)

	privateKey := pem.EncodeToMemory(privateKeyBlock)
	publicKey := pem.EncodeToMemory(publicKeyBlock)

	fmt.Fprintf(w, "your private key is: %v\n\nand public key is: %v", string(privateKey), string(publicKey))

}
