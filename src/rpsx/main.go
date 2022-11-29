package rpsx

type Float float64
type CharName string

type BasicName string
type Basics map[BasicName] Float

type SlotName string

type OrganName string
type Organs map[OrganName] *Organ
type OrganType string
// Every organ type must have its defined max value function.
type OrganMap map[OrganType] func(*Char) Float

type Skill string

// The main structure to store system description.
// Organs, states, skills etc.
type Implementation struct {
	OrganMap OrganMap
}

type Char struct {
	Impl *Implementation
	Name CharName
	Basics Basics
	Organs Organs
}

type Organ struct {
	Type OrganType
	Health Float
}

func (c *Char)Basic(name BasicName) Float {
	return c.Basics[name]
}

func (c Char)Organ(name )
