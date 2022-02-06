package collect

type bareCollection[T []E, E any] struct {
	z any
}

func (c *bareCollection[T, E]) All() T {
	return c.z.(T)
}

func (c *bareCollection[T, E]) Len() int {
	return len(c.z.(T))
}
