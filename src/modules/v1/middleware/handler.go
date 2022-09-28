package middleware

import (
	"fmt"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type Chain []Middleware

func HandlerChain(m ...Middleware) Chain {
	var slice Chain
	return append(slice, m...)
}

func (c Chain) Then(lastHandler http.HandlerFunc) http.HandlerFunc {
	//**expected mid1(mid2(mid3(handlerFunc)))
	for i := range c {
		lastHandler = c[len(c)-1-i](lastHandler)
	}
	return lastHandler
}

func Hello(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		next.ServeHTTP(w, r)
	}
}
