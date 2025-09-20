package main

import (
	"errors"
	"fmt"
)

var ERR_OUT_OF_TEA = errors.New("no more tea available!!")
var ERR_POWER_OUTAGE = errors.New("no electricity to boil water!!")

func makeTea(arg int) error {
	if arg == 0 {
		return nil
	} else if arg == 1 {
		return ERR_OUT_OF_TEA
	} else if arg == 2 {
		return ERR_POWER_OUTAGE
	} else {
		return errors.New("couldn't make tea for some reason!!")
	}
}
func main() {
	fmt.Println("Errors")

	for i := range 4 {
		err := makeTea(i)
		if err == nil {
			fmt.Println("preparing tea, will be delivered shortly")
		} else if errors.Is(err, ERR_OUT_OF_TEA) {
			fmt.Println("we should buy more tea!!")
		} else if errors.Is(err, ERR_POWER_OUTAGE) {
			fmt.Println("check the power box!!")
		} else {
			fmt.Printf("unknown error: %s\n", err)
		}
	}
}
