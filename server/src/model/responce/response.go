package responce

import "cmp_lab/src/structs/opt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response[T any] struct {
	Result opt.Option[T] `json:"result"`
	Error  Error         `json:"error"`
}
