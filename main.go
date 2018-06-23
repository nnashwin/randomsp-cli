package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {
	errCol := color.New(color.FgRed).SprintFunc()
	respCol := color.New(color.FgMagenta).SprintFunc()

	resp, err := StartCli(os.Args)
	if err != nil {
		fmt.Println(errCol(fmt.Sprintf("%s", err)))
		return
	}

	// Print all output to the response
	for _, str := range resp {
		fmt.Printf("\n%s", respCol(fmt.Sprintf(str)))
	}
}
