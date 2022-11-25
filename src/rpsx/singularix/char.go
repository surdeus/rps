package singularix

import (
	. "github.com/surdeus/rps/src/rpsx"
)

const (
)

func NewHumanChar() *Char {
	c := &Char{}
	c.Basics = &Valuers {
		C: c,
		VS :map[string] Valuer {
		"str" : &Strength{NewBasic()},
		"per" : &Perception{NewBasic()},
		"end" : &Endurance{NewBasic()},
		"cha" : &Charisma{NewBasic()},
		"int" : &Intelligence{NewBasic()},
		"agi" : &Agility{NewBasic()},
		"lck" : &Luck{NewBasic()},
		},
	}

	c.Healths = &Valuers {
		C: c,
		VS : map[string] Valuer {
		"rhand" : &HumanHand{NewBasic()},
		"lhand" : &HumanHand{NewBasic()},
		},
	}

	return c
}
