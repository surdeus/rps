package main

import (
	"github.com/surdeus/rps/src/rpsx"
	"fmt"
	//"encoding/json"
	//"github.com/surdeus/rps/src/rpsx"
)

func main() {
	c := rpsx.NewHuman("Jay")
	fmt.Println(c.Organ("rhand"))
	fmt.Println(c.Organ("lleg"))
	fmt.Println(c.Health("lleg"))
	c.AddHealth("lleg", -10000)
	fmt.Println(c.Health("lleg"))
}
