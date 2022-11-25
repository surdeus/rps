package rpsx

import (
	"math"
	"errors"
	"fmt"
)

type Float float32
type Valuers struct {
	VS map[string] Valuer
	C *Char
}

type Char struct {
	Name string
	Basics *Valuers
	Healths *Valuers
	States *Valuers
	Skills *Valuers
}

type Basic struct {
	value Float
}

const (
	Infinity Float = math.MaxFloat32
)

func (vs Valuers) Get(name string) Float {
	ret, ok := vs.VS[name]
	if !ok {
		panic(errors.New("no such valuer"))
	}
	return ret.Get(vs.C)
}

func (vs Valuers) Set(name string, v Float) {
	valuer, ok := vs.VS[name]
	if !ok {
		panic(errors.New("no such valuer"))
	}

	valuer.Set(vs.C, v)
}

func (vs Valuers) Max(name string) Float {
	valuer, ok := vs.VS[name]
	if !ok {
		panic(errors.New("no such valuer"))
	}
	return valuer.Max(vs.C)
}

func (vs Valuers) Min(name string) Float {
	valuer, ok := vs.VS[name]
	if !ok {
		panic(errors.New("no such valuer"))
	}
	return valuer.Min(vs.C)
}

func NewBasic() *Basic {
	return &Basic{}
}

func (b *Basic)Get(c *Char) Float {
	return b.value
}

func (b *Basic)Set(c *Char, f Float) {
	fmt.Println("before:", b.value)
	b.value = Clutch(f, b.Min(c), b.Max(c))
	fmt.Println("after:", b.value)
}

func (b *Basic)Add(c *Char, f Float) {
	b.Set(c, b.Get(c) + f)
}

func (b *Basic)Max(c *Char) Float {
	return Infinity
}

func (b *Basic)Min(c *Char) Float {
	return -Infinity
}

func Clutch(v, min, max Float) Float {
	fmt.Println("v:", v)
	fmt.Println("min:", min)
	fmt.Println("max:", max)
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

type Valuer interface {
	Get(*Char) Float
	Max(*Char) Float
	Min(*Char) Float
	Set(*Char, Float)
	Add(*Char, Float)
}

