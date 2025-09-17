package Handlers

import (
	"crypto_currency/pkg/BlockChain"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type SignUp struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}
type Response struct {
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
	Status  int         `json:"status"`
}
type Keyes struct {
	Public  string `json:"public"`
	Private string `json:"private"`
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

	response := &Response{}
	response.Message = "its your keyes!"
	response.Status = http.StatusOK
	response.Value = Keyes{
		Public: strings.Trim(string(publicKey), "\n"),
		Private: strings.Trim(string(privateKey), "\n"),
	}


	res ,_:= json.MarshalIndent(response, " ", "\n")
	_, err = fmt.Fprint(w, string(res))

	if err != nil {
		log.Printf("we have an error while showing response")
	}
}
