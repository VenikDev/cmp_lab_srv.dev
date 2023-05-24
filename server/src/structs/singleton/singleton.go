package singleton

import "sync"

type Singleton[T any] struct {
	instance *T
	once     sync.Once
}

func (s *Singleton[T]) GetInstance(newInstanceFunc func() T) *T {
	s.once.Do(func() {
		instance := newInstanceFunc()
		s.instance = &instance
	})
	return s.instance
}
