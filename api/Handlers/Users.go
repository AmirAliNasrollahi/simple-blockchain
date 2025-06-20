package Handlers

import (
	"fmt"
	"net/http"
)

func TEST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "now you are in User handler")
}

