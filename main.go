package main

import (
	"fmt"
	"github.com/ru-lai/randomsp"
)

func main() {
	stock := randomsp.GetRandomSPStock()
	fmt.Println(stock)
}
