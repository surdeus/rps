package main

import (
	sinx "github.com/surdeus/rps/src/rpsx/singularix"
	"fmt"
	. "github.com/surdeus/rps/src/rpsx"
)

func main() {
	char := sinx.NewHumanChar()
	char.Basics.Set("str", 7)
	char.Basics.Set("end", 5)
	fmt.Println(char.Basics.Get("str"))
	//fmt.Println(char.Healths.Max("rhand"))
	char.Healths.Set("rhand", Infinity)
	fmt.Println(char.Healths.Get("rhand"))
	fmt.Println(char)
}
