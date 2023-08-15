package checkerror

import (
	"log"
	"os"
)

func CheckError(e error) {
	if e != nil {
		log.Printf("%+v\n", e)
		os.Exit(1)
	}
}
