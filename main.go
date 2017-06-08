package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	b, err := ioutil.ReadFile("./config.hcl")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	config, err := hclParser(string(b))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(config.Option, config.Instance, config.Dns)

	os.Exit(0)
}
