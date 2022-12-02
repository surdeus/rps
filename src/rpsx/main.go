package rpsx

import (
	"math"
)

type Float float64
type CharName string

type BasicName string
type Basics map[BasicName] Float

type SlotName string

type OrganName string
type Organs map[OrganName] *Organ
type OrganKind string
type OrganMap map[OrganKind] func(*Char) Float

type Skill string

type Char struct {
	Name CharName
	Basics Basics
	Organs Organs
}

type Organ struct {
	Kind OrganKind
	Health Float
}

const (
	Infinity Float = math.MaxFloat64
)

func (c *Char)B(name BasicName) Float {
	return c.Basics[name]
}

func (c *Char)O(name OrganName) Float {
	return MainOrganMap[c.Organs[name].Kind](c)
}

func (c *Char)H(name OrganName) Float {
	return c.Organs[name].Health
}

func (c *Char)SetHealth(name OrganName, h Float) {
	c.Organs[name].Health = Clutch(
		h,
		-c.O(name),
		c.O(name),
	)
}

func (c *Char)AddHealth(name OrganName, a Float) {
	c.SetHealth(name, c.H(name) + a)
}

func Clutch(v, min, max Float) Float {
	if v < min {
		return min
	} else if v > max {
		return max
	}

	return v
}

