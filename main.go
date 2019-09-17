package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"golang.org/x/crypto/bcrypt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 3; i++ {
		bcrypt.GenerateFromPassword([]byte("PASSWORD"), bcrypt.DefaultCost)
	}
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
