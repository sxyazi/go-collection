package collect

import (
	"constraints"
)

type numberCollection[T []E, E constraints.Integer | constraints.Float] struct {
	sliceCollection[T, E]
}

func Number[T []E, E constraints.Integer | constraints.Float](items T) *numberCollection[T, E] {
	return &numberCollection[T, E]{sliceCollection[T, E]{Items: items}}
}

func (c *numberCollection[T, E]) Sum() (total E) {
	for _, value := range c.Items {
		total += E(value)
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

	min := c.Items[0]
	for _, value := range c.Items {
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

	max := c.Items[0]
	for _, value := range c.Items {
		if max < value {
			max = value
		}
	}

	return max
}
