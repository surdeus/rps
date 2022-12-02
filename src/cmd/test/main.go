package main

import (
	"github.com/surdeus/rps/src/rpsx"
	"fmt"
	//"encoding/json"
	//"github.com/surdeus/rps/src/rpsx"
)

func main() {
	c := rpsx.NewHuman("Jay")
	fmt.Println(c.O("rhand"))
	fmt.Println(c.O("lleg"))
	fmt.Println(c.H("lleg"))
	c.AddHealth("lleg", -10000)
	fmt.Println(c.H("lleg"))
}
