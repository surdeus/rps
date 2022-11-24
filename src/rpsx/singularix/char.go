package singularix

import (
	. "github.com/surdeus/rps/src/rpsx"
)

func NewChar() *Char {
	c := &Char{}
	c.Basics = Valuers {
		"str" : Strength{NewBasic()},
	}

	str, _ := c.Basics["str"]
	str.Set(7)

	return c
}
