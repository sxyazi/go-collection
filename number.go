package collect

import (
	"constraints"
)

type numberCollection[T []E, E constraints.Integer | constraints.Float] struct {
	sliceCollection[T, E]
}

func Number[T []E, E constraints.Integer | constraints.Float](items T) *numberCollection[T, E] {
	return &numberCollection[T, E]{sliceCollection[T, E]{z: items}}
}

func (c *numberCollection[T, E]) Sum() (total E) {
	for _, value := range c.Items() {
		total += value
	}
	return
}

func (c *numberCollection[T, E]) Avg() E {
	if c.Empty() {
		return 0
	}

	return c.Sum() / E(c.Len())
}

func (c *numberCollection[T, E]) Min() E {
	if c.Empty() {
		return 0
	}

	z := c.Items()
	min := z[0]
	for _, value := range z {
		if min > value {
			min = value
		}
	}

	return min
}

func (c *numberCollection[T, E]) Max() E {
	if c.Empty() {
		return 0
	}

	z := c.Items()
	max := z[0]
	for _, value := range z {
		if max < value {
			max = value
		}
	}

	return max
}
