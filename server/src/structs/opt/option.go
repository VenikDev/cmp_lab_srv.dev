package opt

type Option[T any] struct {
	Value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{&value}
}

func None[T any]() Option[T] {
	return Option[T]{nil}
}

func (o Option[T]) IsSome() bool {
	return o.Value != nil
}

func (o Option[T]) IsNone() bool {
	return o.Value == nil
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("cannot unwrap None")
	}
	return *o.Value
}

func (o Option[T]) ValueOrCall(callback func()) T {
	if o.IsNone() {
		callback()
	}
	return *o.Value
}

func (o Option[T]) ValueOr(value T) T {
	if o.IsNone() {
		return value
	}
	return *o.Value
}

func (o Option[T]) ProcessIfHas(fn func(*T)) T {
	fn(o.Value)
	return *o.Value
}
