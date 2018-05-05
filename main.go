package main

import (
	"fmt"
	"os"
)

func main() {
	err := StartCli(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
