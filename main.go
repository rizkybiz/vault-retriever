package main

import (
	"fmt"
	"os"

	"github.com/rizkybiz/vault-retriever/retriever"
)

// TODO: Implement TLS
//       Implement Wrapping token?

func main() {
	r, err := retriever.New()
	if err != nil {
		fmt.Println("app encountered a problem, exiting.")
		os.Exit(1)
	}
	err = r.Run()
	if err != nil {
		fmt.Println("app encountered a problem, exiting.")
		os.Exit(1)
	}
}
