package main

import (
	"net/http"

	libhttp "github.com/yudai2929/http/pkg"
)

func main() {

	if false {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})

		http.ListenAndServe(":8081", nil)
	}

	err := libhttp.ListenAndServe(8080)

	if err != nil {
		panic(err)
	}

}
