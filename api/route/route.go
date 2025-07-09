package route

import (
	"crypto_currency/api/Handlers"
	"crypto_currency/api/middleware"
	"net/http"
)
// what's different between HandleFunc and HandlerFunc
// first you set Handler with request pattern & second you just set Handler
func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	helloWorldWithMiddleWare := middleware.AuthorizationMiddleware(http.HandlerFunc(Handlers.HelloWorld))

	mux.Handle("/", helloWorldWithMiddleWare)
	return mux
}

// what is different between Handle and HandleFunc ...
// first you should pass the Handler to that
// second just pass the function to that (خودش تبدلیش میکنه به handler)
