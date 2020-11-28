package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/syumai/uuidgenseeded/uuidseeded"
)

var (
	lower = flag.Bool("lower", false, "output result as lower case")
	l = flag.Bool("l", false, "output result as lower case (shorthand of lower)")
)

func realMain(seedReader io.Reader, lower bool) (string, error) {
	// generate UUID
	idStr, err := uuidseeded.New(seedReader)
	if err != nil {
		return "", err
	}

	// switch output case (lower / upper)
	if lower {
		return idStr, nil
	}
	return strings.ToUpper(idStr), nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	// initialize seed reader
	var seedReader io.Reader
	if len(args) > 0 {
		seedReader = strings.NewReader(args[0])
	} else {
		seedReader = os.Stdin
	}

	id, err := realMain(seedReader, *lower || *l)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
