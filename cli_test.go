package main

import (
	"fmt"
	"testing"
)

func TestStartCli(t *testing.T) {
	ran, err := StartCli([]string{"./rsp", "random"})
	if len(ran) < 1 && err != nil {
		fmt.Println("The random method was not able to find a stock.  You might not be connected to the internet.  Please connect and try again")
		t.Fail()
	}

	nas, err := StartCli([]string{"./rsp", "index", "nas"})
	if len(nas) < 1 && err != nil {
		fmt.Println("The nasdaq index method was unable to find a stock.  You might not be connected to the internet.  Please connect and try again")
		t.Fail()
	}

	sp, err := StartCli([]string{"./rsp", "index", "sp"})
	if len(sp) < 1 && err != nil {
		fmt.Println("The sp index method was unable to find a stock.  You might not be connected to the internet.  Please connect and try again")
		t.Fail()
	}

	dax, err := StartCli([]string{"./rsp", "index", "dax"})
	if len(dax) < 1 && err != nil {
		fmt.Println("The dax index method was unable to find a stock.  You might not be connected to the internet.  Please connect and try again")
		t.Fail()
	}

	ft, err := StartCli([]string{"./rsp", "index", "ft"})
	if len(ft) < 1 && err != nil {
		fmt.Println("The ft index method was unable to find a stock.  You might not be connected to the internet.  Please connect and try again")
		t.Fail()
	}

	ift, err := StartCli([]string{"./rsp", "index", "ift"})
	if len(ift) < 1 && err != nil {
		fmt.Println("The ift index method was unable to find a stock.  You might not be connected to the internet.  Please connect and try again")
		t.Fail()
	}
}
