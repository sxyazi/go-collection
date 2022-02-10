package collect

import (
	"golang.org/x/exp/constraints"
)

type NumberCollection[T ~[]E, E constraints.Integer | constraints.Float] struct {
	*SliceCollection[T, E]
}

func UseNumber[T ~[]E, E constraints.Integer | constraints.Float](items T) *NumberCollection[T, E] {
	return &NumberCollection[T, E]{UseSlice[T, E](items)}
}

func (n *NumberCollection[T, E]) Sum() (total E) {
	return Sum[T, E](n.All())
}

func (n *NumberCollection[T, E]) Min() E {
	return Min[T, E](n.All())
}

func (n *NumberCollection[T, E]) Max() E {
	return Max[T, E](n.All())
}

func (n *NumberCollection[T, E]) Sort() *NumberCollection[T, E] {
	n.z = Sort[T, E](n.All())
	return n
}

func (n *NumberCollection[T, E]) SortDesc() *NumberCollection[T, E] {
	n.z = SortDesc[T, E](n.All())
	return n
}

func (n *NumberCollection[T, E]) Avg() float64 {
	return Avg[T, E](n.All())
}

func (n *NumberCollection[T, E]) Median() float64 {
	return Median[T, E](n.All())
}
