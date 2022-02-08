package collect

import (
	"golang.org/x/exp/constraints"
)

type numberCollection[T ~[]E, E constraints.Integer | constraints.Float] struct {
	*SliceCollection[T, E]
}

func UseNumber[T ~[]E, E constraints.Integer | constraints.Float](items T) *numberCollection[T, E] {
	return &numberCollection[T, E]{UseSlice[T, E](items)}
}

func (c *numberCollection[T, E]) Sum() (total E) {
	return Sum[T, E](c.All())
}

func (c *numberCollection[T, E]) Avg() E {
	return Avg[T, E](c.All())
}

func (c *numberCollection[T, E]) Min() E {
	return Min[T, E](c.All())
}

func (c *numberCollection[T, E]) Max() E {
	return Max[T, E](c.All())
}
