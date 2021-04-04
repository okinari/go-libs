package golibs

import (
	"fmt"
	"log"
	"os"
)

func FailOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
		fmt.Fprintf(os.Stderr, "%s\n", err)
		panic(err)
	}
}
