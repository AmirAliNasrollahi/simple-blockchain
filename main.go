package main

import (
	"crypto_currency/api/route"
	"crypto_currency/internal"
	"fmt"
	"log"
	"net/http"
	"time"
)
var Configs []internal.ENV = internal.ReadConf()

func main() {

	// set a multiplaxer for if serve many server in other ports
	mux := route.Routes()	

	addr := Configs[0].ConfValue+":"+Configs[1].ConfValue
	server := &http.Server{
		Addr: addr,
		ReadTimeout: time.Second * 10,
		Handler: mux,
	}

	fmt.Println("the first Server is Running ...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("we have problem on serving")
	}
}
