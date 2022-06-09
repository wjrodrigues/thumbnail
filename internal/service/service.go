package service

type Response[T any] struct {
	result T
	errors error
}

type Service[T any, R any] interface {
	Call(T) Response[R]
}
