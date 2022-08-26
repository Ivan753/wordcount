package main

import (
    "fmt"
    "strings"
    "flag"
    "os"
)

func main() {
    src := readInput()

    words := strings.Fields(src)

    fmt.Println(len(words))
}

func readInput() (src string) {
	flag.Parse()
	src = strings.Join(flag.Args(), " ")
	return src
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
