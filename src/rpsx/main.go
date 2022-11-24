package rpsx

import (
	"math"
)

type Float float32
type Valuers map[string] Valuer

type Char struct {
	Name string
	Basics Valuers
	Healths Valuers
	States Valuers
	Skills Valuers
}

type Basic struct {
	value Float
}

const (
	Infinity Float = math.MaxFloat32
)

func (v Valuers) Get(name string) Float {
	ret, _ := v[name]
	return ret.Get()
}

func NewBasic() *Basic {
	return &Basic{}
}

func (b *Basic)Get() Float {
	return b.value
}

func (b *Basic)Set(f Float) {
	b.value = Clutch(f, b.Min(), b.Max())
}

func (b *Basic)Add(f Float) {
	b.Set(b.Get() + f)
}

func (b *Basic)Max() Float {
	return Infinity
}

func (b *Basic)Min() Float {
	return -Infinity
}

func Clutch(v, min, max Float) Float {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

type Valuer interface {
	Get() Float
	Max() Float
	Min() Float
	Set(Float)
	Add(Float)
}

