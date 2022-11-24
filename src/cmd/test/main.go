package main

import (
	sinx "github.com/surdeus/rps/src/rpsx/singularix"
	"fmt"
)

func main() {
	char := sinx.NewChar()
	fmt.Println(char.Basics.Get("str"))
}
