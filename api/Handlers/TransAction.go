package Handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


func SendCoin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("\n we have an error when reading input\n")
		return
	}
	fmt.Fprintf(w, string(body));
	err = json.Unmarshal(body, Tr)
}
