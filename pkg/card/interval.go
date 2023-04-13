package card

import (
	"errors"
	"math/rand"
)

type Interval struct {
	min, max int
}

var ErrMinMax = errors.New("минимум превышает максимум")

func NewInterval(min, max int) *Interval {
	return &Interval{
		min: min,
		max: max,
	}
}

func (i *Interval) GetMin() int {
	return i.min
}

func (i *Interval) SetMin(v int) error {
	if v > i.max {
		return ErrMinMax
	}

	i.min = v

	return nil
}

func (i *Interval) GetMax() int {
	return i.max
}

func (i *Interval) SetMax(v int) error {
	if v < i.min {
		return ErrMinMax
	}

	i.max = v

	return nil
}

/*
Возвращает случайное значение в пределах интервала
*/
func (i *Interval) RandomValue() (v int) {
	d := i.GetMax() - i.GetMin()
	if d > 0 {
		v = rand.Intn(d)
	}

	v += i.GetMin()

	return
}
