package main

import (
	"fmt"

	"github.com/captain-corgi/go-mod-example/greet"
)

func main() {
	myName := "Anh Le"
	fmt.Println(greet.HelloHuman(myName))
	fmt.Println(greet.HelloPredator(myName))
}
