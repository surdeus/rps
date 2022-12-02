package rpsx

// Singularix setting implementation

var (
	MainOrganMap = OrganMap{
		"human-hand" : HumanHand,
		"human-leg" : HumanLeg,
	}
)

func NO(kind OrganKind) *Organ {
	return &Organ{kind, 0}
}

func NewHuman(name CharName) *Char {
	c := &Char{
		Name : name,
	}

	c.Basics = Basics{
		"str" : 5,
		"end" : 5,
	}

	c.Organs = Organs {
		"rhand" : NO("human-hand"),
		"lhand" : NO("human-hand"),
		"rleg" : NO("human-leg"),
		"lleg" : NO("human-leg"),
	}

	for v := range c.Organs {
		c.SetHealth(v, Infinity)
	}

	return c
}

func HumanHand(c *Char) Float {
	return c.B("end") * 4
}

func HumanLeg(c *Char) Float {
	return c.B("end") * 5
}

