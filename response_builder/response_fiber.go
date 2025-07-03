package responsebuilder

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseFiber[T any] struct {
	context  *fiber.Ctx
	response *ResponseJson[T]
}

func NewResponseFiber[T any](context *fiber.Ctx) *ResponseFiber[T] {
	return &ResponseFiber[T]{
		context:  context,
		response: NewResponseJson[T](),
	}
}

func (rf *ResponseFiber[T]) Status(status int) *ResponseFiber[T] {
	rf.response.SetStatus(status)
	return rf
}

func (rf *ResponseFiber[T]) Message(message string) *ResponseFiber[T] {
	rf.response.SetMessage(message)
	return rf
}

func (rf *ResponseFiber[T]) Data(data *T) *ResponseFiber[T] {
	rf.response.SetData(data)
	return rf
}

func (rf *ResponseFiber[T]) Pagination(totalRows int, page int, size int) *ResponseFiber[T] {
	rf.response.SetPagination(totalRows, page, size)
	return rf
}

func (rf *ResponseFiber[T]) Errors(errors *T) *ResponseFiber[T] {
	rf.response.SetErrors(errors)
	return rf
}

func (rf *ResponseFiber[T]) Build() error {
	return rf.context.Status(rf.response.Status).JSON(rf.response)
}
