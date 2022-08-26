package main

import (
    "fmt"
    "strings"
    "flag"
    "errors"
    "os"
)

func main() {
    src, err := readInput()

    if err != nil {
        fail(err)
    }

    words := strings.Fields(src)

    fmt.Println(len(words))
}

func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), " ")
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
