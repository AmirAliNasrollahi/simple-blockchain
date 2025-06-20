package main

import (
	"crypto_currency/api/Handlers"
	"crypto_currency/internal"
	"fmt"
	"log"
	"net/http"
	"time"
)
var Configs []internal.ENV = internal.ReadConf()

func main() {

	addr := Configs[0].ConfValue+":"+Configs[1].ConfValue
	server := &http.Server{
		Addr:addr,
		ReadTimeout: time.Second * 10,
		Handler: nil,
	}
	
	http.HandleFunc("/", Handlers.TEST)
	fmt.Println("the Server is Running ...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("we have problem on serving")
	}
}
