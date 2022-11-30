package rpsx

// Singularix setting implementation

var (
	MainOrganMap = OrganMap{
		"human-hand" : HumanHand,
		"human-leg" : HumanLeg,
	}
)

func NewOrgan(kind OrganKind) *Organ {
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
		"rhand" : NewOrgan("human-hand"),
		"lhand" : NewOrgan("human-hand"),
		"rleg" : NewOrgan("human-leg"),
		"lleg" : NewOrgan("human-leg"),
	}

	for v := range c.Organs {
		c.SetHealth(v, Infinity)
	}

	return c
}

func HumanHand(c *Char) Float {
	return c.Basic("end") * 4
}

func HumanLeg(c *Char) Float {
	return c.Basic("end") * 5
}

