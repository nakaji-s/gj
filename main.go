package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var err error
	var contents []byte

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		bio := bufio.NewReader(os.Stdin)
		contents, err = bio.ReadBytes(0)
		if err != io.EOF {
			fmt.Printf("File error: %v\n", err)
			os.Exit(1)
		}
	} else {
		filename := args[0]
		contents, err = ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("File error: %v\n", err)
			os.Exit(1)
		}
	}

	var buffer bytes.Buffer
	err = json.Indent(&buffer, contents, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		str, err := buffer.ReadString(0)
		if err != io.EOF {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(str)
	}
}
