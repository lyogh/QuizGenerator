package card

import "errors"

type Interval struct {
	min, max uint
}

var ErrMinMax = errors.New("минимум превышает максимум")

func NewInterval(min, max uint) *Interval {
	return &Interval{
		min: min,
		max: max,
	}
}

func (i *Interval) GetMin() uint {
	return i.min
}

func (i *Interval) SetMin(v uint) error {
	if v > i.max {
		return ErrMinMax
	}

	i.min = v

	return nil
}

func (i *Interval) GetMax() uint {
	return i.max
}

func (i *Interval) SetMax(v uint) error {
	if v < i.min {
		return ErrMinMax
	}

	i.max = v

	return nil
}
