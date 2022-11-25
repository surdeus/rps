package singularix

import (
	. "github.com/surdeus/rps/src/rpsx"
	"fmt"
)

type HumanEye struct {
	*Basic
}

type HumanHand struct {
	*Basic
}

func (h *HumanHand)Max(c *Char) Float {
	fmt.Println("in")
	return c.Basics.Get("end") * 4
}

