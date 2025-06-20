package main

import (
	"crypto_currency/api/Handlers"
	"crypto_currency/internal"
	"fmt"
	"log"
	"net/http"
)
var Configs []internal.ENV = internal.ReadConf()

func main() {
	http.HandleFunc("/", Handlers.TEST)

	addr := Configs[0].ConfValue+":"+Configs[1].ConfValue
	fmt.Println("the Server is Running ...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("we have problem on serving")
	}
}
