package main

import (
	sinx "github.com/surdeus/rps/src/rpsx/singularix"
	"fmt"
	"encoding/json"
	//"github.com/surdeus/rps/src/rpsx"
)

func main() {
	c := sinx.NewHuman()
	c.SetBasic("str", 7)
	c.SetBasic("end", 5)
	fmt.Println(c.Basic("str"))
	fmt.Println(c.MaxHealth("reye"))
	js, err := json.Marshal(*c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(js)
	}
}
