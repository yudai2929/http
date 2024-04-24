package main

import (
	"net/http"

	libhttp "github.com/yudai2929/http/pkg"
)

func main() {
	err := libhttp.ListenAndServe(8080)

	if err != nil {
		panic(err)
	}

}
