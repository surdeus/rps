package sinx

import (
	. "github.com/surdeus/rps/src/rpsx"
)

// Singularix setting implementation

var (
	MainOrganMap = OrganMap{
		"human-hand" : HumanHand,
		"human-leg" : HumanLeg,
	}
	Impl = Implementation{
	}
)

func NewHuman(name CharName) *Char {
	c := &Char{
		Impl : Impl,
		Name : name,
	}

	c.Organs = Organs {
	}

}

func HumanHand(c *Char) Float {
	return c.Basic("end") * 4
}

func HumanLeg(c *Char) Float {
	return c.Basic("end") * 5
}
